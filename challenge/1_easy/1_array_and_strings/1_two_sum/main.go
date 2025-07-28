package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var seenNums = make(map[int]int)
	for idx, num := range nums {
		remain := target - num
		id, ok := seenNums[remain]
		if !ok {
			seenNums[num] = idx
			continue
		}

		return []int{id, idx}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}
