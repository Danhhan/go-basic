package main

import "fmt"

func main() {
	var result []int
	arr1:= []int{3,2,4}
	result = twoSum(arr1, 6);
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	total := 0
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			total = nums[i] + nums[j]
			if total == target {
				result = append(result, i,j)
				return result
			}
		}
	}
	return nil 
}
