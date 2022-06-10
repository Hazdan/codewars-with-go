package kata

import "strings"

func ToCamelCase(str string) string {
	size := len(str)
  out := ""
  beforeChar := ""
  for i := 0; i < size; i++ {
    char := string(str[i])
    
    if beforeChar == "-"  || beforeChar == "_" {
      char = strings.ToUpper(char)
    }
    if char != "-"  && char != "_" {
      out += char
    }
    
    beforeChar = char
  }
	return out
}
