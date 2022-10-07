package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day6"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(6)
  fmt.Println(day6.Part1(input))
  fmt.Println(day6.Part2(input))
}
