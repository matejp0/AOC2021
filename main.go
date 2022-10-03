package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day4"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(4)
  fmt.Println(day4.Part1(input))
  fmt.Println(day4.Part2(input))
}
