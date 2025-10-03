# Reverse Integer

Given a signed 32-bit integer `x`, return `x` with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range `[-231, 231 - 1]`, then return `0`.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

# Example

> [!NOTE]     
> Input: x = 123    
> Output: 321       

> [!NOTE]      
> Input: x = -123       
> Output: -321     

> [!NOTE]     
> Input: x = 120    
> Output: 21      




# Constraints

 - `-231 <= x <= 231 - 1`

# Go Code

```go
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
```