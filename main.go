package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day9"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(9)
  fmt.Println(day9.Part1(input))
  fmt.Println(day9.Part2(input))
}
