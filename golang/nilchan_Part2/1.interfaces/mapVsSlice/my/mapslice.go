package main

import (
	"fmt"
	"time"
)

type Item struct {
	ID int
}

type List struct {
	items []Item
}

func (l *List) Add(item Item) {
	l.items = append(l.items, item)
}

type Mapa struct {
	mapa map[int]int
}

func (m *Mapa) GetInfo() {

}

func main() {
	iterations := 1000_000

	// Тестируем Map
	m := make(map[int]int)
	startMap := time.Now()
	for i := 0; i < iterations; i++ {
		m[i] = i
	}
	fmt.Printf("Map:   %v\n", time.Since(startMap))

	// Тестируем Slice
	list := List{}
	startSlice := time.Now()
	for i := 0; i < iterations; i++ {
		list.Add(Item{ID: i})
	}
	fmt.Printf("Slice: %v\n", time.Since(startSlice))

	// Поиск в Map: O(1) — мгновенно
	startMapSearch := time.Now()
	for i := 0; i < iterations; i++ {
		m.GetInfo
	}
	fmt.Printf("Map search:   %v\n", time.Since(startMapSearch))

	// Поиск в Slice: O(n) — долго, так как перебираем массив
	startSliceSearch := time.Now()
	for _, v := range list.items {
		if v.ID == i {
			break
		}
	}
	fmt.Printf("Slice search: %v\n", time.Since(startSliceSearch))
}
