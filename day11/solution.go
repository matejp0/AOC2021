package day11

import (
  "strconv"
)

const SIDE = 10

func Part1(input []string) int {
  array := parse(input)
  var flashes int

  for step := 0; step < 100; step++ {
    visited := make(map[[2]int]struct{})
    for y := 0; y < SIDE; y++ {
      for x := 0; x < SIDE; x++ {
        array, visited = process(array, y, x, visited)
        for key, _ := range visited { array[key[0]][key[1]] = 0 }
      }
    }
    flashes += len(visited)

  }

  return flashes
}

func Part2(input []string) int {
  array := parse(input)

  visited := make(map[[2]int]struct{})
  var step int

  for len(visited) != SIDE*SIDE {
    visited = make(map[[2]int]struct{}) // clear map
    step++

    for y := 0; y < SIDE; y++ {
      for x := 0; x < SIDE; x++ {
        array, visited = process(array, y, x, visited)
        for key, _ := range visited { array[key[0]][key[1]] = 0 }
      }
    }

  }

  return step
}

func process(array [SIDE][SIDE]int, y int, x int, flashed map[[2]int]struct{}) ([SIDE][SIDE]int, map[[2]int]struct{}) {
  if y < 0 || y > SIDE - 1 || x < 0 || x > SIDE - 1 {
    return array, flashed
  }

  array[y][x]++
  _, contains := flashed[[...]int{y, x}]
  if array[y][x] > 9 && !contains {
    flashed[[...]int{y, x}] = struct{}{}
    for i := -1; i <= 1; i++ {
      for j := -1; j <= 1; j++ {
        if i != 0 || j != 0 { // isn't itself
          array, flashed = process(array, y+i, x+j, flashed)
        }
      }
    }
  }

  return array, flashed
}

func parse(input []string) [SIDE][SIDE]int {
  var array [SIDE][SIDE]int
  for y, line := range input {
    for x, char := range line {
      num, err :=  strconv.Atoi(string(char))
      if err != nil { panic(err) }
      array[y][x] = num
    }
  }

  return array
}
