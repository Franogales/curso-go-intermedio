package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exist := m.cache[key]
	m.lock.Unlock()
	if !exist {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 42, 38}
	var wg sync.WaitGroup
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("Calculate: %d, Time: %s, Result: %d\n", index, time.Since(start), value)
		}(n)

	}
	wg.Wait()
}
