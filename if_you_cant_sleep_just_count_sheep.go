package kata

import "fmt"

func countSheep(num int) string {
  var out string
  for i := 1; i <= num; i++ {
    out += fmt.Sprintf("%d sheep...", i)
  }
  return out
}
