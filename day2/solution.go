package day2

import (
  "strings"
  "strconv"
)
func Part1(input []string) int {
  var horizontal, depth int
  for _, line := range input {
    command, number := splitLine(line, " ")
    n, err := strconv.Atoi(number)
    if err != nil { panic(err) }
    switch command {
    case "up":
      depth -= n
    case "down":
      depth += n
    case "forward":
      horizontal += n
    }
  }
  return horizontal*depth
}

func Part2(input []string) int {
  var horizontal, depth, aim int
  for _, line := range input {
    command, number := splitLine(line, " ")
    n, err := strconv.Atoi(number)
    if err != nil { panic(err) }
    switch command {
    case "up":
      aim -= n
    case "down":
      aim += n
    case "forward":
      horizontal += n
      depth += aim*n
    }
  }
  return horizontal*depth
}

func splitLine(s string, sep string) (string, string) {
  x := strings.Split(s, sep)
  return x[0], x[1]
}
