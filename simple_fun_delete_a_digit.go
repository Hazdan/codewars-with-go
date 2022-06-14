package kata

import (
  "strconv"
  "sort"
)

func DeleteDigit(n int) int {
  number := strconv.Itoa(n)
  numbers := []int{}
  for i := range number {
    newNumber, _ := strconv.Atoi(number[:i] + number[i+1:])
    numbers = append(numbers, newNumber)
  }
  sort.Ints(numbers)
  return numbers[len(numbers) - 1]
}
