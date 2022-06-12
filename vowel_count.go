package kata

import "regexp"

func GetCount(str string) (count int) {
  re := regexp.MustCompile(`[^aeiou]`)
  
  vowels := re.ReplaceAllString(str, "")
  return len(vowels)
}
