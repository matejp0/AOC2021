package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day5"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(5)
  fmt.Println(day5.Part1(input))
  fmt.Println(day5.Part2(input))
}
