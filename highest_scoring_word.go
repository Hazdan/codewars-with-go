package kata

import "strings"

func High(str string) string {
  words := strings.Split(str, " ")
  wordIndex := 0
  highest := 0
  
  for i, word := range words {
    sum := 0
    for _, char := range word {
      sum += (int(char) - 96)
    }
    if sum > highest {
      highest = sum
      wordIndex = i
    }
  }
  return words[wordIndex]
}
