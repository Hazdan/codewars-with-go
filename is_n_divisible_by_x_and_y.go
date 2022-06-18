package kata

func IsDivisible(n, x, y int) bool {
    if x == 0 || y == 0 || n%x != 0 || n%y != 0 {
    return false
  }
  return true
}
