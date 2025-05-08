package main

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    s := len(nums)
	
    for i := 0; i < s; i++ {
        for j := i + 1; j < s; j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    
	//error handling
    if result != nil {
        fmt.Printf("Indices: %v (nums[%d] + nums[%d] = %d)\n", result, result[0], result[1], target)
    } else {
        fmt.Println("No two sum solution found.")
    }
}
