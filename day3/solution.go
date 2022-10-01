package day3

import (
  "strconv"
  "strings"
)

func Part1(input []string) int {
  length := len(input)
  linelength := len(input[0])
  max := square(2, linelength) - 1
  
  array := make([]int, linelength)
  for _, line := range input {
    for i, r := range line {
      if r == '1' { array[i]++ }
    }
  }

  var sb strings.Builder
  for _, num := range array {
    if num > length/2 {
      sb.WriteRune('1')
    } else {
      sb.WriteRune('0')
    }
  }
  
  str := sb.String()
  g, err := strconv.ParseInt(str, 2, 0)
  if err != nil { panic(err) }

  return int(g)*(max-int(g))
}

func Part2(input []string) int {
  linelength := len(input[0])
  for i := 0; i < linelength; i++ {
    for index, value := range input {
      if value[i] == '1'
    }
  }
}

func square(a, b int) int {
  var result = a
  for i := 1; i < b; i++ {
    result = result * a
  }

  return result
}
