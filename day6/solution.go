package day6

import (
  "strings"
  "strconv"
)

func Part1(input []string) uint64 {
  return do(input, 80)

}

func Part2(input []string) uint64 {
  return do(input, 256)
}


func do(input []string, days uint16) uint64{
  var count uint64
  cache := make(map[uint8]uint64, 0)
  for _, v := range parse(input) {
    if cache[v] != 0 {
      count += cache[v]
    } else {
      lant := countLanternfish(v, 1, days)
      cache[v] = lant
      count += lant 
    }
  }
  return count
}

func countLanternfish(timer uint8, day uint16, days uint16) uint64 {
  var count uint64 = 1

  for i := day; i <= days; i++ {
    if timer == 0 {
      timer = 6
      count += countLanternfish(9, i, days)
    } else {
      timer--
    }
  }

  return count
}

func parse(input []string) []uint8 {
  list := make([]uint8, 0)
  for _, v := range strings.Split(input[0], ",") {
    n, err := strconv.Atoi(v)
    if err != nil { panic(err) }
    list = append(list, uint8(n))
  }
  return list
}
