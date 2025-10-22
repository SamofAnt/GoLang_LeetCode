package Easy

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
