package kata

import "strings"

func Map(f func(s string) string, slice []string) []string {
	out := []string{}
	for _, v := range slice {
		out = append(out, f(v))
	}
	return out
}

func SpinWords(str string) string {
  reverse := func(word string) string {
    reversed := word
    if len(word) >= 5 {
      reversed = ""
      for _, char := range word {
        reversed = string(char) + reversed
      }
    }
    return reversed
  }
  
  return strings.Join(Map(reverse, strings.Split(str," ")), " ")
}
