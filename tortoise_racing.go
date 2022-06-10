package kata

const secondsPerHour = 3600
const secondsPerMinute = 60
const minutesPerHour = 60

func Race(v1, v2, g int) [3]int {
  out := [3]int{-1, -1, -1}
  if v1 < v2 {
    speedDifference := v2 - v1
    secondsToCatch := secondsPerHour * g / speedDifference
    
    out[0] = secondsToCatch / secondsPerHour
    out[1] = (secondsToCatch % secondsPerHour) / minutesPerHour
    out[2] = secondsToCatch % secondsPerMinute
  }
  return out
}
