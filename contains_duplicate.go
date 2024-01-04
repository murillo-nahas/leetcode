package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for _, num := range nums {
		if set[num] {
			return true
		} else {
			set[num] = true
		}
	}
	return false
}
