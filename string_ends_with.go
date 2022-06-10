package kata

import "strings"

func solution(str, ending string) bool {
	str = strings.Trim(str, " ")
	ending = strings.Trim(ending, " ")
  from := len(str) - len(ending)
	if from < 0 {
		from = 0
	}
	return str[from:] == ending
}
