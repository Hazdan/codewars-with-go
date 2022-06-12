package kata

import "sort"

func InAscOrder(numbers []int) bool {
  copyNumbers := make([]int, len(numbers))
  copy(copyNumbers, numbers)
  sort.Ints(copyNumbers)
  out := true
  for i, n := range numbers {
    if copyNumbers[i] != n {
      out = false
      break
    }
  } 
  return out
}
