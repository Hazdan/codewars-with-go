package kata

import "fmt"

func CreatePhoneNumber(arr [10]uint) string {
  numbers := make([]interface{}, len(arr))
  for i, v := range arr {
    numbers[i] = v
  }
  return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", numbers...)
}
