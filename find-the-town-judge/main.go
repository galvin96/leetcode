package main

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	exists := struct{}{}
	trusts := make(map[int]int)
	untrusts := make(map[int]struct{})
	for i := 0; i < len(trust); i++ {
		untrusts[trust[i][0]] = exists
		if _, ok := untrusts[trust[i][1]]; ok == false {
			trusts[trust[i][1]]++
		}
		if _, ok := trusts[trust[i][0]]; ok {
			delete(trusts, trust[i][0])
		}
	}
	if len(trusts) == 1 {
		for k := range trusts {
			if trusts[k] == n-1 {
				return k
			}
		}
	}
	return -1
}

func main() {
	findJudge(2, [][]int{{1, 2}})
	findJudge(3, [][]int{{1, 2}, {2, 3}})
	findJudge(3, [][]int{{1, 3}, {2, 3}})
	findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})
	findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}})
}
