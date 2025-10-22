package Easy

func IsValid(s string) bool {
	stackBrackets := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stackBrackets = append(stackBrackets, char)
		} else {
			if len(stackBrackets) == 0 || stackBrackets[len(stackBrackets)-1] != pairs[char] {
				return false
			}
			stackBrackets = stackBrackets[:len(stackBrackets)-1]
		}
	}
	return len(stackBrackets) == 0
}
