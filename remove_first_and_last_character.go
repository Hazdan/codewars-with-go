package kata

func RemoveChar(word string) string {
  word = word[1:len(word)-1]
  return word
}
