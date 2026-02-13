// Package main da.
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	data map[string]int
	mu   sync.RWMutex
}

// Set устанавливает значение в мапу.
func (s *SafeMap) Set(key string, val int) {
	// Только одна горутина может писать
	// и читать, другие ждут
	s.mu.Lock()
	defer s.mu.Unlock()

	// Пишем в мапу данные
	// Они запишутся так как мы пишем
	// прям в нее а не в копию
	// key - string. val - int
	// если бы писали в в более простом виде,
	// то: m["da"] = 2
	s.data[key] = val
}

// Get получает результат, было ли это в мапе.
func (s *SafeMap) Get(key string) (int, bool) {
	// Может много кто читать, но пишет всегда один
	// Так будет гарантия что прочитающая горутина
	// получит точно актуальные данные

	// Несколько читателей могут работать одновременно, но
	// если появился писатель, то читатели ждут
	s.mu.RLock()
	defer s.mu.RUnlock()

	// В val вернется int
	val, ok := s.data[key]

	return val, ok
}

func v1() {
	start := time.Now()
	wg := sync.WaitGroup{}
	// Сразу инициализируем мапу чтобы в нее можно
	// можно было что-то положить
	da := &SafeMap{data: make(map[string]int)}

	// Читают
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			val, ok := da.Get("da")
			if ok {
				fmt.Printf("Читатель %d: ключ 'da', значение %d\n", id, val)
			}
		}(i)
	}

	// Пишут
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			da.Set("da", id*10)
		}(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

// Этот вариант гарантирует что порядок выполнения кода
// будет один и тот же. Всегда выполнится сначала
// писатели, потом читатели.
func v2() {
	start := time.Now()
	wg := sync.WaitGroup{}
	da := &SafeMap{data: make(map[string]int)}

	// сигнал о готовности всего
	ready := make(chan struct{})

	// Тут нет смысла делать так, можно сделать так
	// da.Set("da", 123)
	// При этом не создавая канал и не закрывая его
	// все бы работало так же

	// Писатель
	go func() {
		da.Set("da", 123)
		close(ready)
	}()

	// Код блокируется на этом месте, он
	// ждет закрытия или изменение канала
	<-ready

	// Читатели
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			val, ok := da.Get("da")
			if ok {
				fmt.Printf("Читатель %d: ключ=da значение %d\n ", id, val)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func v3() {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}
	da := &SafeMap{data: make(map[string]int)}

	for i := 0; i < 500; i++ {
		wg2.Add(1)

		go func() {
			defer wg2.Done()
			da.Set("da", 123)
		}()
	}
	wg2.Wait()

	// Читатели
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			val, ok := da.Get("da")
			if ok {
				fmt.Printf("Читатель %d: ключ=da значение %d\n ", id, val)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}

func main() {
	// v1()
	// v2()
	v3() // третий мне нравится больше всего
}
