package day7

import (
  "strings"
  "strconv"
  "math"
)

func Part1(input []string) int {
  numbers, biggest := parse(input)

  lowest := math.MaxInt64
  for position := 0; position <= biggest; position++ {
    var fuel int
    for _, n := range numbers {
      fuel += abs(n - position)
    }
    if fuel < lowest {
      lowest = fuel
    }
  }

  return lowest
}

func Part2(input []string) int {
  numbers, biggest := parse(input)

  lowest := math.MaxInt64
  for position := 0; position <= biggest; position++ {
    var fuel int
    for _, n := range numbers {
      for i := 1; i <= abs(n - position); i++ {
        fuel += i
      }
    }
    if fuel < lowest {
      lowest = fuel
    }
  }

  return lowest
}

func parse(input []string) ([]int, int) {
  s := strings.Split(input[0], ",")
  var numbers = make([]int, len(s))
  var biggest int
  for i, v := range strings.Split(input[0], ",") {
    number, err := strconv.Atoi(v)
    if err != nil { panic(err) }

    numbers[i] = number
    if number > biggest { biggest = number }
  }

  return numbers, biggest
}

func abs(number int) int {
  if number < 0 { number = -number }
  return number
}
