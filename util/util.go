package util

func Min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}

	return res
}

func Max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if res < nums[i] {
			res = nums[i]
		}
	}

	return res
}
