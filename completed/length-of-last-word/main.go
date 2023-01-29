package main

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	i := len(s) - 1
	for s[i] == ' ' {
		i--
	}
	ret := 0
	for i > -1 && s[i] != ' ' {
		ret++
		i--
	}
	return ret
}

func main() {
	lengthOfLastWord("")
	lengthOfLastWord("a")
	lengthOfLastWord("Hello World")
	lengthOfLastWord("   fly me   to   the moon  ")
	lengthOfLastWord("luffy is still joyboy")
}
