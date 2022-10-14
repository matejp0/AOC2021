package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day10"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(10)
  fmt.Println(day10.Part1(input))
  fmt.Println(day10.Part2(input))
}
