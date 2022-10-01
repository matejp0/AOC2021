package day3

import (
  "strconv"
  "strings"
)

func Part1(input []string) int {
  var gammarate, epsilonrate strings.Builder

  for i := 0; i < len(input[0]); i++ {
    most, least := mostCommon(input, i)
    gammarate.WriteRune(most)
    epsilonrate.WriteRune(least)
  }
  
  gamma, er := strconv.ParseInt(gammarate.String(), 2, 0)
  if er != nil { panic(er) }

  epsilon, err := strconv.ParseInt(epsilonrate.String(), 2, 0)
  if err != nil { panic(err) }

  return int(gamma*epsilon)
}


func mostCommon(lines []string, i int) (rune, rune) {
  var count int
  var length int = len(lines)
  for _, line := range lines {
     if line[i] == '1' { count++ }
  }

  if count > length/2 {
    return '1', '0'
  }
  return '0', '1'
}
