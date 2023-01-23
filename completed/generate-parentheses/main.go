package main

import "strings"

func backtrack(S []string, left int, right int, n int, ret *[]string) {
	if len(S) == 2*n {
		*ret = append(*ret, strings.Join(S, ""))
		return
	} else {
		if right < left {
			S = append(S, ")")
			right++
			backtrack(S, left, right, n, ret)
			S = S[:len(S)-1]
			right--
		}
		if left < n {
			S = append(S, "(")
			left++
			backtrack(S, left, right, n, ret)
		}

	}
}

func generateParenthesis(n int) []string {
	var ret []string

	backtrack([]string{}, 0, 0, n, &ret)

	return ret
}

func main() {
	generateParenthesis(3)
	generateParenthesis(2)
}
