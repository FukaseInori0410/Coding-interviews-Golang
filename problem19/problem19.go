package problem19

func Match(str, pattern string) bool {
	if len(pattern) == 0 {
		return len(str) == 0
	}
	var first bool
	if len(str) > 0 && (str[0] == pattern[0] || pattern[0] == '.') {
		first = true
	}
	if len(pattern) > 1 && pattern[1] == '*' {
		return first && Match(str[1:], pattern) || Match(str, pattern[2:])
	}
	return first && Match(str[1:], pattern[1:])
}
