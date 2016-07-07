package main

import (
	"fmt"
	"math"
)

func main() {
	names := []string{"Orlando", "Miami", "Tampa", "Jacksonville"}
	for _, name := range names {
		printName(name)
	}
}

func printName(n string) {
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println(n)
}
