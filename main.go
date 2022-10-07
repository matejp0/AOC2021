package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day7"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(7)
  fmt.Println(day7.Part1(input))
  fmt.Println(day7.Part2(input))
}
