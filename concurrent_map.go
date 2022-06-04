package cmap

import (
	"encoding/json"
	"sync"
)

var ShardCount = 32

// ConcurrentMap A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (ShardCount) map shards.
type ConcurrentMap[V any] []*ConcurrentMapShared[V]

// ConcurrentMapShared A "thread" safe string to anything map.
type ConcurrentMapShared[V any] struct {
	items        map[string]V
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

// New Creates a new concurrent map.
func New[V any]() ConcurrentMap[V] {
	m := make(ConcurrentMap[V], ShardCount)
	for i := 0; i < ShardCount; i++ {
		m[i] = &ConcurrentMapShared[V]{items: make(map[string]V)}
	}
	return m
}

// GetShard returns shard under given key
func (m ConcurrentMap[V]) GetShard(key string) *ConcurrentMapShared[V] {
	return m[uint(fnv32(key))%uint(ShardCount)]
}

func (m ConcurrentMap[V]) MSet(data map[string]V) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items[key] = value
		shard.Unlock()
	}
}

// Set Sets the given value under the specified key.
func (m ConcurrentMap[V]) Set(key string, value V) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// UpsertCb Callback to return new element to be inserted into the map
// It is called while lock is held, therefore it MUST NOT
// try to access other keys in same map, as it can lead to deadlock since
// Go sync.RWLock is not reentrant
type UpsertCb[V any] func(exist bool, valueInMap V, newValue V) V

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCb
func (m ConcurrentMap[V]) Upsert(key string, value V, cb UpsertCb[V]) (res V) {
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items[key]
	res = cb(ok, v, value)
	shard.items[key] = res
	shard.Unlock()
	return res
}

// SetIfAbsent Sets the given value under the specified key if no value was associated with it.
func (m ConcurrentMap[V]) SetIfAbsent(key string, value V) bool {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	_, ok := shard.items[key]
	if !ok {
		shard.items[key] = value
	}
	shard.Unlock()
	return !ok
}

// Get retrieves an element from map under given key.
func (m ConcurrentMap[V]) Get(key string) (V, bool) {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// SetOrGet Sets the given value under the specified key if no value was associated with it, otherwise return existing value
func (m ConcurrentMap[V]) SetOrGet(key string, value V) (actual V, absent bool) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	actual, ok := shard.items[key]
	if !ok {
		shard.items[key] = value
		actual = value
	}
	shard.Unlock()
	return actual, !ok
}

// GetValue Get retrieves an element from map under given key.
func (m ConcurrentMap[V]) GetValue(key string) (val V) {
	val, _ = m.Get(key)
	return
}

// Count returns the number of elements within the map.
func (m ConcurrentMap[V]) Count() int {
	count := 0
	for i := 0; i < ShardCount; i++ {
		shard := m[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Has Looks up an item under specified key
func (m ConcurrentMap[V]) Has(key string) bool {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// See if element is within shard.
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}

// Remove removes an element from the map.
func (m ConcurrentMap[V]) Remove(key string) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// RemoveCb is a callback executed in a map.RemoveCb() call, while Lock is held
// If returns true, the element will be removed from the map
type RemoveCb[V any] func(key string, v V, exists bool) bool

// RemoveCb locks the shard containing the key, retrieves its current value and calls the callback with those params
// If callback returns true and element exists, it will remove it from the map
// Returns the value returned by the callback (even if element was not present in the map)
func (m ConcurrentMap[V]) RemoveCb(key string, cb RemoveCb[V]) bool {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items[key]
	remove := cb(key, v, ok)
	if remove && ok {
		delete(shard.items, key)
	}
	shard.Unlock()
	return remove
}

// Pop removes an element from the map and returns it
func (m ConcurrentMap[V]) Pop(key string) (v V, exists bool) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	v, exists = shard.items[key]
	delete(shard.items, key)
	shard.Unlock()
	return v, exists
}

// IsEmpty checks if map is empty.
func (m ConcurrentMap[V]) IsEmpty() bool {
	return m.Count() == 0
}

// Tuple Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type Tuple[V any] struct {
	Key string
	Val V
}

// Iter returns an iterator which could be used in a for range loop.
//
// Deprecated: using IterBuffered() will get a better performence
func (m ConcurrentMap[V]) Iter() <-chan Tuple[V] {
	chans := snapshot(m)
	ch := make(chan Tuple[V])
	go fanIn(chans, ch)
	return ch
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m ConcurrentMap[V]) IterBuffered() <-chan Tuple[V] {
	chans := snapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan Tuple[V], total)
	go fanIn(chans, ch)
	return ch
}

// Clear removes all items from map.
func (m ConcurrentMap[V]) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a snapshot of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func snapshot[V any](m ConcurrentMap[V]) (chans []chan Tuple[V]) {
	// When you access map items before initializing.
	if len(m) == 0 {
		panic(`cmap.ConcurrentMap is not initialized. Should run New() before usage.`)
	}
	chans = make([]chan Tuple[V], ShardCount)
	wg := sync.WaitGroup{}
	wg.Add(ShardCount)
	// Foreach shard.
	for index, shard := range m {
		go func(index int, shard *ConcurrentMapShared[V]) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan Tuple[V], len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- Tuple[V]{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// fanIn reads elements from channels `chans` into channel `out`
func fanIn[V any](chans []chan Tuple[V], out chan Tuple[V]) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan Tuple[V]) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}

// Items returns all items as map[string]V
func (m ConcurrentMap[V]) Items() map[string]V {
	tmp := make(map[string]V)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}

	return tmp
}

// IterCb Iterator callbacalled for every key,value found in
// maps. RLock is held for all calls for a given shard
// therefore callback sess consistent view of a shard,
// but not across the shards
type IterCb[V any] func(key string, v V)

// IterCb Callback based iterator, cheapest way to read
// all elements in a map.
func (m ConcurrentMap[V]) IterCb(fn IterCb[V]) {
	for idx := range m {
		shard := (m)[idx]
		shard.RLock()
		for key, value := range shard.items {
			fn(key, value)
		}
		shard.RUnlock()
	}
}

// Keys returns all keys as []string
func (m ConcurrentMap[V]) Keys() []string {
	count := m.Count()
	ch := make(chan string, count)
	go func() {
		// Foreach shard.
		wg := sync.WaitGroup{}
		wg.Add(ShardCount)
		for _, shard := range m {
			go func(shard *ConcurrentMapShared[V]) {
				// Foreach key, value pair.
				shard.RLock()
				for key := range shard.items {
					ch <- key
				}
				shard.RUnlock()
				wg.Done()
			}(shard)
		}
		wg.Wait()
		close(ch)
	}()

	// Generate keys
	keys := make([]string, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}

// MarshalJSON Reviles ConcurrentMap "private" variables to json marshal.
func (m ConcurrentMap[V]) MarshalJSON() ([]byte, error) {
	// Create a temporary map, which will hold all item spread across shards.
	tmp := make(map[string]V)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}
	return json.Marshal(tmp)
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}
