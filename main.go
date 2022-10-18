package main

import (
  "fmt"
  "github.com/matejp0/aoc2021/day13"
  "github.com/matejp0/aoc2021/readinput"
)

func main() {
  var input = readinput.Read(13)
  fmt.Println(day13.Part1(input))
   fmt.Println(day13.Part2(input))
}
