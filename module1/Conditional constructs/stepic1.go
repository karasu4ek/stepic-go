package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	var firstHalf, secondHalf int
	for i := 0; i < 6; i++ {
		if i < 3 {
			firstHalf += num % 10
		} else {
			secondHalf += num % 10
		}
		num = num / 10
	}
	if firstHalf == secondHalf {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
