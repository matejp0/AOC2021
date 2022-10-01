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
  
  gamma := binToDec(gammarate.String())
  epsilon := binToDec(epsilonrate.String())

  return int(gamma*epsilon)
}

func Part2(input []string) int {
  generator  := binToDec(findOne(input, 0, true))
  scrubber := binToDec(findOne(input, 0, false))

  return int(generator*scrubber)
}

func findOne(lines []string, index int, mostIsCommon bool) string {
  if len(lines) == 1 { return lines[0] }

  most, least := mostCommon(lines, index)

  criterium := least
  if mostIsCommon { criterium = most }
  
  filtered := make([]string, 0)
  for _, value := range lines {
    if rune(value[index]) == criterium {
      filtered = append(filtered, value)
    }
  }

  return findOne(filtered, index+1, mostIsCommon)
}

func mostCommon(lines []string, i int) (rune, rune) {
  var count int
  var length int = len(lines)
  for _, line := range lines {
     if line[i] == '1' { count++ }
  }

  if float64(count) >= float64(length)/2 {
    return '1', '0'
  }
  return '0', '1'
}

func binToDec(s string) int {
  i, err := strconv.ParseInt(s, 2, 0)
  if err != nil {
    panic(err)
  }
  return int(i)
}
