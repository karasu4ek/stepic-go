package main

import "fmt"

func main() {
	sum := 100
	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}
	fmt.Println(sum)
}
