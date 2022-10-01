package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day1"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(1)
  fmt.Println(day1.Part1(input))
  fmt.Println(day1.Part2(input))
}
