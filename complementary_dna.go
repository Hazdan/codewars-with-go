package kata

func DNAStrand(dna string) string {
  out := ""
  for _, char := range dna {
    switch char {
      case 'A':
        out += "T"
      case 'T':
        out += "A"
      case 'C':
        out += "G"
      case 'G':
        out += "C"
    }
  }
  return out
}
