package main

import (
	"fmt"
	"math"
	"sync"
)

func isPrime(n int) bool {
	if n == 0 || n == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	numGoroutines := 5
	var result []int
	input := make([]int, 0)
	for i := n; i <= m; i++ {
		input = append(input, i)
	}
	chunkSize := len(input) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if start >= len(input) { // Handle cases where numGoroutines > len(input)
			break
		}
		if end > len(input) {
			end = len(input)
		}
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()
			localResult := []int{}
			for _, num := range chunk {
				if isPrime(num) {
					localResult = append(localResult, num)
				}
			}
			mu.Lock()
			result = append(result, localResult...)
			mu.Unlock()
		}(input[start:end])
	}
	wg.Wait()

	fmt.Print(result)
}
