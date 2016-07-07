package main

import "fmt"

func main() {
	names := []string{"Orlando", "Miami", "Tampa", "Jacksonville"}
	for _, name := range names {
		printName(name)
	}
}

func printName(n string) {
	fmt.Println(n)
}
