package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day11"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(11)
  fmt.Println(day11.Part1(input))
  fmt.Println(day11.Part2(input))
}
