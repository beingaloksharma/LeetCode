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
