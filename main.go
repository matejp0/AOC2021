package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day12"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(12)
  fmt.Println(day12.Part1(input))
  fmt.Println(day12.Part2(input))
}
