package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/puzzle1"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(1)
  fmt.Println(puzzle1.Part1(input))
  fmt.Println(puzzle1.Part2(input))
}
