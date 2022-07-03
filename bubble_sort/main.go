package main

import "fmt"

func bubble_sort(nums []int32) []int32 {

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
	var array []int32 = []int32{8,5,3,4,1,7,2,6,0}
	var sorted []int32 = bubble_sort(array)
	fmt.Println("Sorted: ", sorted)

}