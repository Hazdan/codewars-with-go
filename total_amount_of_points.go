package kata

import "strconv"

func Points(games []string) int {
  out := 0
  for _, game := range games {
    x, _ := strconv.Atoi(string(game[0]))
    y, _ := strconv.Atoi(string(game[2]))
    if x > y {
      out += 3
    } else if x == y {
      out += 1
    }
  }
  return out
}
