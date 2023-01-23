package main

func compare(char1, char2 byte) bool {
	return char1 == '(' && char2 == ')' || (char1 == '{' && char2 == '}') || (char1 == '[' && char2 == ']')
}

func isValid(s string) bool {
	if len(s) == 2 {
		if compare(s[0], s[1]) {
			return true
		} else {
			return false
		}
	} else if len(s)%2 == 1 {
		return false
	}

	i := len(s) - 1
	for i > 0 {
		if !(compare(s[0], s[i])) {
			i--
		} else if !isValid(s[1:i]) {
			/**
			* Validate inside
			* if string inside is wrong, decrease i to continue
			**/
			i--
		} else {
			break
		}
	}

	pass := true

	/**
	* Validate right side
	**/
	if i+1 < len(s) && !isValid(s[i+1:]) {
		pass = false
	}

	return pass
}
