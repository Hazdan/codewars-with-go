package kata


func FakeBin(x string) string {
  bin := ""
  for _, char := range x {
    if char < 53 {
      bin += "0"
    } else {
      bin += "1"
    }
  }
  return bin
}
