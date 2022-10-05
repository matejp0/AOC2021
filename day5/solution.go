package day5

import (
  "strings"
  "strconv"
)

func Part1(input []string) int{
  points := parseInput(input, false)

  var count int
  for _, point := range points {
    if point >= 2 {
      count++
    }
  }
  return count
}

func Part2(input []string) int{
  points := parseInput(input, true)

  var count int
  for _, point := range points {
    if point >= 2 {
      count++
    }
  }
  return count
}

func parseInput(s []string, diagonals bool) map[[2]int]int {
  points := make(map[[2]int]int, 0)
  for _, line := range s {
    arr := strings.Split(line, " -> ")
    start := splitPair(arr[0])
    end := splitPair(arr[1])

    if isDiagonal(start, end) && diagonals {
      start, end = sortArr(start, end, 0)
      multiplier := 1
      if(start[1] > end[1]) {
        multiplier = -1
      }
      for i := 0; i <= end[0]-start[0]; i++ {
        points[[2]int{start[0]+i, start[1]+multiplier*i}]++
      }
    } else if start[0] == end[0] {
      start, end = sortArr(start, end, 1)
      for y := start[1]; y <= end[1]; y++ {
        points[[2]int{start[0], y}]++
      }
    } else if start[1] == end[1] {
      start, end = sortArr(start, end, 0)
      for x := start[0]; x <= end[0]; x++ {
        points[[2]int{x, start[1]}]++
      }
    }
  }

  return points
}

func isDiagonal(start [2]int, end [2]int) bool {
  return abs(end[0] - start[0]) == abs(end[1] - start[1])
}

func abs(i int) int{
  if i < 0 {
    return -i
  }
  return i
}

func splitPair(s string) [2]int {
  arr := strings.Split(s, ",")

  x, errx := strconv.Atoi(arr[0])
  if errx != nil { panic(errx) }

  y, erry := strconv.Atoi(arr[1])
  if erry != nil { panic(erry) }

  return [2]int{x, y}
}

func sortArr(start [2]int, end [2]int, i byte) ([2]int, [2]int) {
  if start[i] > end [i] {
    return end, start
  }
  return start, end
}
