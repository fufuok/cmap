//go:build go1.18
// +build go1.18

package cmap

import (
	"encoding/json"
	"sync"

	"github.com/fufuok/cmap/internal/xxhash"
)

// MapOf A "thread" safe map of type string:Anything.
// To avoid lock bottlenecks this map is dived to several (ShardCount) map shards.
type MapOf[K comparable, V any] struct {
	shards   []*SharedOf[K, V]
	sharding func(key K) uint64
}

// SharedOf A "thread" safe string to anything map.
type SharedOf[K comparable, V any] struct {
	items        map[K]V
	sync.RWMutex // Read Write mutex, guards access to internal map.
}

func create[K comparable, V any](sharding func(key K) uint64) *MapOf[K, V] {
	m := &MapOf[K, V]{
		sharding: sharding,
		shards:   make([]*SharedOf[K, V], ShardCount),
	}
	for i := 0; i < ShardCount; i++ {
		m.shards[i] = &SharedOf[K, V]{items: make(map[K]V)}
	}
	return m
}

// NewOf Creates a new concurrent map.
func NewOf[K comparable, V any]() *MapOf[K, V] {
	return NewWithCustomShardingFunction[K, V](xxhash.GenHasher64[K]())
}

// NewWithCustomShardingFunction Creates a new concurrent map.
func NewWithCustomShardingFunction[K comparable, V any](sharding func(key K) uint64) *MapOf[K, V] {
	return create[K, V](sharding)
}

// GetShard returns shard under given key
func (m *MapOf[K, V]) GetShard(key K) *SharedOf[K, V] {
	return m.shards[uint(m.sharding(key))%uint(ShardCount)]
}

func (m *MapOf[K, V]) MSet(data map[K]V) {
	for key, value := range data {
		shard := m.GetShard(key)
		shard.Lock()
		shard.items[key] = value
		shard.Unlock()
	}
}

// Set Sets the given value under the specified key.
func (m *MapOf[K, V]) Set(key K, value V) {
	// Get map shard.
	shard := m.GetShard(key)
	shard.Lock()
	shard.items[key] = value
	shard.Unlock()
}

// UpsertCbOf Callback to return new element to be inserted into the map
// It is called while lock is held, therefore it MUST NOT
// try to access other keys in same map, as it can lead to deadlock since
// Go sync.RWLock is not reentrant
type UpsertCbOf[V any] func(exist bool, valueInMap V, newValue V) V

// Upsert Insert or Update - updates existing element or inserts a new one using UpsertCbOf
func (m *MapOf[K, V]) Upsert(key K, value V, cb UpsertCbOf[V]) (res V) {
	shard := m.GetShard(key)
	shard.Lock()
	v, ok := shard.items[key]
	res = cb(ok, v, value)
	shard.items[key] = res
	shard.Unlock()
	return res
}

// SetIfAbsent Sets the given value under the specified key if no value was associated with it.
func (m *MapOf[K, V]) SetIfAbsent(key K, value V) bool {
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
func (m *MapOf[K, V]) Get(key K) (V, bool) {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// Get item from shard.
	val, ok := shard.items[key]
	shard.RUnlock()
	return val, ok
}

// Count returns the number of elements within the map.
func (m *MapOf[K, V]) Count() int {
	count := 0
	for i := 0; i < ShardCount; i++ {
		shard := m.shards[i]
		shard.RLock()
		count += len(shard.items)
		shard.RUnlock()
	}
	return count
}

// Has Looks up an item under specified key
func (m *MapOf[K, V]) Has(key K) bool {
	// Get shard
	shard := m.GetShard(key)
	shard.RLock()
	// See if element is within shard.
	_, ok := shard.items[key]
	shard.RUnlock()
	return ok
}

// Remove removes an element from the map.
func (m *MapOf[K, V]) Remove(key K) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	delete(shard.items, key)
	shard.Unlock()
}

// RemoveCbOf is a callback executed in a map.RemoveCbOf() call, while Lock is held
// If returns true, the element will be removed from the map
type RemoveCbOf[K any, V any] func(key K, v V, exists bool) bool

// RemoveCb locks the shard containing the key, retrieves its current value and calls the callback with those params
// If callback returns true and element exists, it will remove it from the map
// Returns the value returned by the callback (even if element was not present in the map)
func (m *MapOf[K, V]) RemoveCb(key K, cb RemoveCbOf[K, V]) bool {
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
func (m *MapOf[K, V]) Pop(key K) (v V, exists bool) {
	// Try to get shard.
	shard := m.GetShard(key)
	shard.Lock()
	v, exists = shard.items[key]
	delete(shard.items, key)
	shard.Unlock()
	return v, exists
}

// IsEmpty checks if map is empty.
func (m *MapOf[K, V]) IsEmpty() bool {
	return m.Count() == 0
}

// TupleOf Used by the Iter & IterBuffered functions to wrap two variables together over a channel,
type TupleOf[K comparable, V any] struct {
	Key K
	Val V
}

// IterBuffered returns a buffered iterator which could be used in a for range loop.
func (m *MapOf[K, V]) IterBuffered() <-chan TupleOf[K, V] {
	chans := snapshotOf(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan TupleOf[K, V], total)
	go fanInOf(chans, ch)
	return ch
}

// Clear removes all items from map.
func (m *MapOf[K, V]) Clear() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// Returns a array of channels that contains elements in each shard,
// which likely takes a snapshotOf of `m`.
// It returns once the size of each buffered channel is determined,
// before all the channels are populated using goroutines.
func snapshotOf[K comparable, V any](m *MapOf[K, V]) (chans []chan TupleOf[K, V]) {
	// When you access map items before initializing.
	if len(m.shards) == 0 {
		panic(`cmap.MapOf is not initialized. Should run New() before usage.`)
	}
	chans = make([]chan TupleOf[K, V], ShardCount)
	wg := sync.WaitGroup{}
	wg.Add(ShardCount)
	// Foreach shard.
	for index, shard := range m.shards {
		go func(index int, shard *SharedOf[K, V]) {
			// Foreach key, value pair.
			shard.RLock()
			chans[index] = make(chan TupleOf[K, V], len(shard.items))
			wg.Done()
			for key, val := range shard.items {
				chans[index] <- TupleOf[K, V]{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// fanInOf reads elements from channels `chans` into channel `out`
func fanInOf[K comparable, V any](chans []chan TupleOf[K, V], out chan TupleOf[K, V]) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan TupleOf[K, V]) {
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
func (m *MapOf[K, V]) Items() map[K]V {
	tmp := make(map[K]V)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}

	return tmp
}

// IterCbOf Iterator callbacalled for every key,value found in
// maps. RLock is held for all calls for a given shard
// therefore callback sess consistent view of a shard,
// but not across the shards
type IterCbOf[K comparable, V any] func(key K, v V)

// IterCb Callback based iterator, cheapest way to read
// all elements in a map.
func (m *MapOf[K, V]) IterCb(fn IterCbOf[K, V]) {
	for idx := range m.shards {
		shard := (m.shards)[idx]
		shard.RLock()
		for key, value := range shard.items {
			fn(key, value)
		}
		shard.RUnlock()
	}
}

// Keys returns all keys as []string
func (m *MapOf[K, V]) Keys() []K {
	count := m.Count()
	ch := make(chan K, count)
	go func() {
		// Foreach shard.
		wg := sync.WaitGroup{}
		wg.Add(ShardCount)
		for _, shard := range m.shards {
			go func(shard *SharedOf[K, V]) {
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
	keys := make([]K, 0, count)
	for k := range ch {
		keys = append(keys, k)
	}
	return keys
}

// MarshalJSON Reviles MapOf "private" variables to json marshal.
func (m *MapOf[K, V]) MarshalJSON() ([]byte, error) {
	// Create a temporary map, which will hold all item spread across shards.
	tmp := make(map[K]V)

	// Insert items to temporary map.
	for item := range m.IterBuffered() {
		tmp[item.Key] = item.Val
	}
	return json.Marshal(tmp)
}

// UnmarshalJSON Reverse process of Marshal.
func (m *MapOf[K, V]) UnmarshalJSON(b []byte) (err error) {
	tmp := make(map[K]V)

	// Unmarshal into a single map.
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	// foreach key,value pair in temporary map insert into our concurrent map.
	for key, val := range tmp {
		m.Set(key, val)
	}
	return nil
}
