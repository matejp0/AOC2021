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
  
  queue := make([][2]int, 0)
  visited := make(map[[2]int]struct{}, 0)
  basins := make([]int, 0)

  for _, riskPoint := range riskPoints {
    queue = append(queue, riskPoint)

    for len(queue) != 0 {

      var elem [2]int = queue[0]
      _, ok := visited[elem]
      if !ok {
        visited[elem] = struct{}{}

        queue, _ = dequeue(queue)
        
        if elem[1] < lenx-1 && points[elem[0]][elem[1]+1] < 9 {
          loc := [...]int{elem[0], elem[1]+1}
          queue = append(queue, loc)
        }
        if elem[1] > 0 && points[elem[0]][elem[1]-1] < 9 {
          loc := [...]int{elem[0], elem[1]-1}
          queue = append(queue, loc)
        }
        if elem[0] < leny-1 && points[elem[0]+1][elem[1]] < 9 {
          loc := [...]int{elem[0]+1, elem[1]}
          queue = append(queue, loc)
        }
        if elem[0] > 0 && points[elem[0]-1][elem[1]] < 9 {
          loc := [...]int{elem[0]-1, elem[1]}
          queue = append(queue, loc)
        }

      } else {
        if len(queue) > 0{
          queue, _ = dequeue(queue)
        } else {
          break
        }
      }
    }
    queue = nil
    basins = append(basins, len(visited))
    visited = make(map[[2]int]struct{}, 0)
  }

  sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })
  
  return basins[0]*basins[1]*basins[2]
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
      points[y][x] = toInt(byte(char))
    }
  }

  return points
}
func toInt(char byte) int {
  num, err := strconv.Atoi(string(char))
  if err != nil {
    panic(err)
  }

  return num
}
