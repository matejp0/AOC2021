package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day3"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(3)
  fmt.Println(day3.Part1(input))
  fmt.Println(day3.Part2(input))
}
