package kata

import "strings"

func ReverseWords(str string) string {
  var items []string = strings.Split(str, " ")
  var reverse_items []string
  for i := len(items) - 1; i >= 0; i-- {
    reverse_items = append(reverse_items, items[i])
  }
  return strings.Join(reverse_items, " ")
}
