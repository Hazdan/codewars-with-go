package kata

func TwiceAsOld(dadYearsOld, sonYearsOld int) int { 
  sonDoubleAge := sonYearsOld * 2
  difference := sonDoubleAge - dadYearsOld
  if difference < 0 {
    difference *= -1
  }
  return difference;
}
