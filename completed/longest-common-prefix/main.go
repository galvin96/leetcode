package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	sample := strs[0]
	i := 0
	for i < len(sample) {
		j := 1
		for ; j < len(strs); j++ {
			if i > len(strs[j])-1 || strs[j][i] != sample[i] {
				break
			}
		}
		if j != len(strs) {
			return sample[0:i]
		}
		i++
		if i == len(sample) {
			return sample[0:i]
		}
	}

	return ""
}

func main() {
	longestCommonPrefix([]string{"flower", "flower", "flower", "flower"})
	longestCommonPrefix([]string{"cir", "car"})
	longestCommonPrefix([]string{"aaa"})
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
}
