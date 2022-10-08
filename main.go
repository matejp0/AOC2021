package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day8"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(8)
  fmt.Println(day8.Part1(input))
  fmt.Println(day8.Part2(input))
}
