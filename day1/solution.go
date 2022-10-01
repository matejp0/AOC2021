package day1

import (
  "strconv"
)

func Part1(input []string) int {
  last := 0
  count := 0
  for _, value := range input {
    number := atoi(value)
    if number > last { count++ }
    last = number
  }
  return count-1
}

func Part2(input []string) int {
  last := 0
  count := 0
  for i := 0; i < len(input)-2; i++ {
    sum := atoi(input[i]) + atoi(input[i+1]) + atoi(input[i+2])
    if sum > last { count++ }
    last = sum
  }
  
  return count-1
}

func atoi(s string) int {
  i, err := strconv.Atoi(s)
  if err != nil { panic(err) }
  return i
}
