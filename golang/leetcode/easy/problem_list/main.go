package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{1, 2, 3, 4, 9}, 13))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, ok := m[complement]; ok {
			return []int{index, i}
		}
		m[num] = i
	}
	return nil
}
