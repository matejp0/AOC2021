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
  leny := len(points)
  lenx := len(points[0])
  for y, line := range points {
    for x, num := range line {
      if x < lenx-1 && points[y][x+1] <= num { continue }
      if x > 0 && points[y][x-1] <= num { continue }
      if y < leny-1 && points[y+1][x] <= num { continue }
      if y > 0 && points[y-1][x] <= num { continue }

      riskPoints = append(riskPoints, [...]int{y, x})
    }
  }
  
  basins := make([]int, 0)

  for _, riskPoint := range riskPoints {
    visited := make(map[[2]int]struct{}, 0)
    queue := append(make([][2]int, 0), riskPoint)

    for len(queue) != 0 {
      var elem [2]int
      queue, elem = dequeue(queue)
      visited[elem] = struct{}{}
      
      if loc := [...]int{elem[0], elem[1]+1}; elem[1] < lenx-1 && points[loc[0]][loc[1]] < 9 && !contains(visited, loc) {
          queue = append(queue, loc)
      }
      if loc := [...]int{elem[0], elem[1]-1}; elem[1] > 0 && points[loc[0]][loc[1]] < 9 && !contains(visited, loc) {
        queue = append(queue, loc)
      }
      if loc := [...]int{elem[0]+1, elem[1]}; elem[0] < leny-1 && points[loc[0]][loc[1]] < 9 && !contains(visited, loc) {
        queue = append(queue, loc)
      }
      if loc := [...]int{elem[0]-1, elem[1]}; elem[0] > 0 && points[loc[0]][loc[1]] < 9 && !contains(visited, loc) {
        queue = append(queue, loc)
      }
    }

    basins = append(basins, len(visited))
  }

  sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })
  
  return basins[0]*basins[1]*basins[2]
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
