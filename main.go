package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	names := []string{"Orlando", "Miami", "Tampa", "Jacksonville"}

	var wg sync.WaitGroup
	wg.Add(len(names))

	for _, name := range names {
		go printName(name, &wg)
	}

	wg.Wait()
}

func printName(n string, wg *sync.WaitGroup) {
	defer wg.Done()

	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println(n)
}
