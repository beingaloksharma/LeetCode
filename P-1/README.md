# 1. Two Sum  

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have `exactly one solution`, and you may not use the same element twice.

> [!NOTE]  
> Example 
> Input: nums = [1,2,0]
> Output: 3
> Explanation: The numbers in the range [1,2] are all in the array.

> [!NOTE] 
> Example 
> Input: nums = [3,4,-1,1]
> Output: 2
> Explanation: 1 is in the array but 2 is missing.

> [!NOTE] 
> Example 
> Input: nums = [7,8,9,11,12]
> Output: 1
> Explanation: The smallest positive integer 1 is missing.

# Constraints

 - `1 <= nums.length <= 105`
 - `-231 <= nums[i] <= 231 - 1`

# Go Code

```go
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
   var result []int
	for i := range nums {
		for j := range nums {
			if i < j {
				if nums[i]+nums[j] == target {
					fmt.Printf("i - %d :: j - %d :: = %d\n", nums[i], nums[j], nums[i]+nums[j])
					result = append(result, i, j)
				}
			}
		}
	}
	return result
}

```