package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day14"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(14)
  fmt.Println(day14.Part1(input))
  fmt.Println(day14.Part2(input))
}
