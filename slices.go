package main

import "fmt"

func main() {
	prime := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = prime[1:4]
	fmt.Println(s)
}
