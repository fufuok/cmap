package cmap_test

import (
	"fmt"

	"github.com/fufuok/cmap"
)

func ExampleMap() {
	m := cmap.New()
	m.Set("A", 1)
	v := m.Upsert("A", 2, func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
		if valueInMap == 1 {
			return newValue.(int) + 1
		}
		return 0
	})
	fmt.Println(v)
	fmt.Println(m.Get("A"))
	m.SetIfAbsent("B", 42)
	m.Remove("A")
	fmt.Println(m.Count())
	for item := range m.IterBuffered() {
		fmt.Println(item)
	}
	m.Clear()

	// Output:
	// 3
	// 3 true
	// 1
	// {B 42}
}
