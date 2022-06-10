package kata

// import (
//   "fmt"
//   "strings"
// )

// func CountBits(n uint) int {
//   return len(strings.Replace(fmt.Sprintf("%b", n), "0", "", -1))
// }

import "math/bits"

func CountBits(n uint) int {
  return bits.OnesCount(n)
}
