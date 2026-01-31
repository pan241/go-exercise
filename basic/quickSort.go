package main

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	target := nums[0]
	l, r := -1, len(nums)
	for l < r {
		for l++; nums[l] < target; l++ {
		}
		for r--; nums[r] > target; r-- {
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	quickSort(nums[:l])
	quickSort(nums[r+1:])
}
