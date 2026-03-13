package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	servers := []string{"SRV-1", "SRV-2", "SRV-3", "SRV-4", "SRV-5"}
	logsChan := make(chan int)
	var wg sync.WaitGroup // 1. Счетчик для горутин

	for _, srv := range servers {
		wg.Add(1)
		go func(name string) {
			defer wg.Done() // Уменьшаем счетчик, когда горутина закончит работу

			iterations := 1 + rand.IntN(5)
			for i := 0; i < iterations; i++ {
				time.Sleep(time.Duration(100+rand.IntN(700)) * time.Millisecond)
				logsChan <- rand.IntN(100)
			}
		}(srv)
	}

	// 2. Отдельная горутина, которая закроет канал, когда ВСЕ воркеры закончат
	go func() {
		wg.Wait()
		close(logsChan)
	}()

	// 3. Главный сборщик с таймером
	timeout := time.After(2 * time.Second) // Канал, который "стрельнет" через 2 сек
	total := 0

loop:
	for {
		select {
		case val, ok := <-logsChan:
			if !ok {
				fmt.Println("Все данные собраны!")
				break loop
			}
			total += val
			fmt.Println("Получено:", val)
		case <-timeout:
			fmt.Println("ВРЕМЯ ВЫШЛО! Собрали что успели.")
			break loop
		}
	}

	fmt.Printf("Итоговый результат: %d\n", total)
}
