package main

import (
	"fmt"
	"log"
	"time"
)

type Memory struct {
	f     Function
	cache map[int]FunctionResult
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
	result, exist := m.cache[key]
	if !exist {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 42, 38}
	for _, n := range fibo {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Calculate: %d, Time: %s, Result: %d\n", n, time.Since(start), value)

	}
}
