package TwoSum

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for i, val := range nums {
		comp := target - val

		if idx, isFound := hashTable[comp]; isFound {
			return []int{i, idx}
		}
		hashTable[val] = i
	}

	return []int{-1, -1}
}
