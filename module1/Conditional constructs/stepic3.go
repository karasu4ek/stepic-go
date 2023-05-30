package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a1 := a / 100
	a2 := a/10 - a1*10
	a3 := a - a1*100 - a2*10
	if a1 == a2 || a2 == a3 || a3 == a1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
