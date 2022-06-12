package kata

func Map(f func(i int, v int) int, slice []int) []int {
  out := []int{}
  for i, v := range slice {
    out = append(out, f(i, v))
  }
  return out
}

func Incrementer(numbers []int) []int {
  add := func(i int, v int) int {
    return ((i + 1) + v) % 10
  }
  return Map(add, numbers)
}
