package kata

func HowMuchILoveYou(i int) string {
  petals := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}
  return petals[(i - 1) % len(petals)]
}
