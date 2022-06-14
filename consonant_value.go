package kata

import "regexp"

func High(words []string) int {
  highest := 0
  for _, word := range words {
    sum := 0
    for _, char := range word {
      sum += (int(char) - 96)
    }
    if sum > highest {
      highest = sum
    }
  }
  return highest
}

func solve(str string) int {
  re := regexp.MustCompile("[aeiou]")
  items := re.Split(str, -1)
  return High(items)
}
