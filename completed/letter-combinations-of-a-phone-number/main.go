package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	keys := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ret := make([]string, 0)

	for _, char := range keys[digits[0]-'0'] {
		ret = append(ret, string(char))
	}

	for i := 1; i < len(digits); i++ {
		digit := digits[i] - '0'
		arr := []string{}
		for k := 0; k < len(keys[digit]); k++ {
			for j := 0; j < len(ret); j++ {
				arr = append(arr, ret[j]+string(keys[digit][k]))
			}
		}
		ret = arr
	}

	return ret
}

func main() {
	letterCombinations("")
	letterCombinations("2")
	letterCombinations("234")
	letterCombinations("23")
}
