package kata

func CountSheeps(numbers []bool) int {
  count := 0
  for _, v := range numbers {
    if v {
      count += 1
    }
  }
  return count
}
