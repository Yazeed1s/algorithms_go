package main

import (
	"fmt"
	"math/rand"
)

func bubble_sort(nums []int) []int {

	var swapped bool = true
	for i := 0; i < len(nums); i++ {
		swapped = !swapped
		for j := 0 ; j < len(nums) - i - 1; j++ {
			if nums[j] > nums[j + 1] {
				temp := nums[j]
				nums[j] = nums[j + 1]
				nums[j + 1] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return nums
}

func main() {
	var array = rand.Perm(100)
	var sorted []int = bubble_sort(array)
	fmt.Println("Sorted: ", sorted)

}