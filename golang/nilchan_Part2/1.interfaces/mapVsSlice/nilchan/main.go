package main

import (
	"fmt"
	"strconv"
	"time"
)

// Описание платежа
type PaymentInfo struct {
	ID          int
	Description string
}

// Модуль со слайсом
type PaymentModuleWithSlice struct {
	s []PaymentInfo
}

func (m *PaymentModuleWithSlice) AddInfo(info PaymentInfo) {
	m.s = append(m.s, info)
}

// Метод поиска в слайсе (перебором)
func (m *PaymentModuleWithSlice) GetInfo(id int) PaymentInfo {
	for _, v := range m.s {
		if v.ID == id {
			return v
		}
	}
	return PaymentInfo{}
}

// Модуль с мапой
type PaymentModuleWithMap struct {
	m map[int]PaymentInfo
}

func (m *PaymentModuleWithMap) AddInfo(info PaymentInfo) {
	m.m[info.ID] = info
}

// Метод поиска в мапе (прямой доступ)
func (m *PaymentModuleWithMap) GetInfo(id int) PaymentInfo {
	return m.m[id]
}

func main() {
	iterations := 100_000 // Количество операций
	pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{m: make(map[int]PaymentInfo)}

	// --- ТЕСТ ДОБАВЛЕНИЯ В СЛАЙС ---
	before := time.Now()
	for i := 0; i < iterations; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}
		pSlice.AddInfo(info)
	}
	fmt.Println("slice add:", time.Since(before))

	// --- ТЕСТ ДОБАВЛЕНИЯ В МАПУ ---
	before = time.Now()
	for i := 0; i < iterations; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}
		pMap.AddInfo(info)
	}
	fmt.Println("map add:  ", time.Since(before))

	fmt.Println("---------------------------------")

	// --- ТЕСТ ПОИСКА В СЛАЙСЕ ---
	// Это самый медленный замер, так как GetInfo делает цикл внутри цикла
	before = time.Now()
	for i := 0; i < iterations; i++ {
		info := pSlice.GetInfo(i)
		_ = info
	}
	fmt.Println("slice get:", time.Since(before))

	// --- ТЕСТ ПОИСКА В МАПЕ ---
	before = time.Now()
	for i := 0; i < iterations; i++ {
		info := pMap.GetInfo(i)
		_ = info
	}
	fmt.Println("map get:  ", time.Since(before))
}
