package day9

import (
  "strconv"
//  "fmt"
  "sort"
)

func Part1(input []string) int {
  points := parse(input)
  var riskSum int
  for y, line := range points {
    for x, num := range line {
      if x < len(line)-1 && points[y][x+1] <= num { continue }
      if x > 0 && points[y][x-1] <= num { continue }
      if y < len(points)-1 && points[y+1][x] <= num { continue }
      if y > 0 && points[y-1][x] <= num { continue }

      riskSum += num + 1
    }
  }
  return riskSum
}

func Part2(input []string) int {
  points := parse(input)
  riskPoints := make([][2]int, 0)
  for y, line := range points {
    for x, num := range line {
      if x < len(line)-1 && points[y][x+1] <= num { continue }
      if x > 0 && points[y][x-1] <= num { continue }
      if y < len(points)-1 && points[y+1][x] <= num { continue }
      if y > 0 && points[y-1][x] <= num { continue }

      riskPoints = append(riskPoints, [...]int{y, x})
    }
  }

  basins := make([]int, 0)
  for _, riskPoint := range riskPoints {
    basins = append(basins, explore(riskPoint[0], riskPoint[1], points, make(map[[2]int]struct{}, 0)))
  }
  sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })

  return basins[0]*basins[1]*basins[2]
}

func explore(y int, x int, points [][]int, visited map[[2]int]struct{}) int {
  boundsY := 0 <= y && y < len(points)
  boundsX := 0 <= x && x < len(points[0])
  if !boundsX || !boundsY { return 0 }

  if points[y][x] == 9 { return 0 }

  _, v := visited[[...]int{y, x}]
  if v { return 0 }
  visited[[...]int{y, x}] = struct{}{}

  count := 1
  
  count += explore(y-1, x, points, visited)
  count += explore(y, x+1, points, visited)
  count += explore(y+1, x, points, visited)
  count += explore(y, x-1, points, visited)

  return count
}

func contains(visited map[[2]int]struct{}, point [2]int) bool {
  _, ok := visited[point]
  return ok
}
func dequeue(queue [][2]int) ([][2]int, [2]int) {
  if len(queue) == 1 {
    return nil, queue[0]
  }
  return queue[1:], queue[0]
}

func parse(input []string) [][]int {
  points := make([][]int, len(input))
  for y, line := range input {
    points[y] = make([]int, len(line))
    for x, char := range line {
      num, err := strconv.Atoi(string(char))
      if err != nil { panic(err) }
      points[y][x] = num
    }
  }

  return points
}
