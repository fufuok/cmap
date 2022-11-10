//go:build go1.18
// +build go1.18

package cmap_test

import (
	"fmt"

	"github.com/fufuok/cmap"
)

func ExampleNewOf() {
	m := cmap.NewOf[string, int]()
	m.Set("A", 1)
	v := m.Upsert("A", 2, func(exist bool, valueInMap int, newValue int) int {
		if valueInMap == 1 {
			return newValue + 1
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
