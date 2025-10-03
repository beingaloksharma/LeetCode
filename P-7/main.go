package main

import "fmt"

func main() {
	x := 123
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var rev int
	limit := 2147483648
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
	}
	fmt.Println("Reverse is :: ", rev)
	if rev > limit || rev < -limit {
		return 0
	}
	return rev
}
