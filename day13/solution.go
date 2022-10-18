package day13

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
  points, folds := parse(input) 

  folds = folds[:1]

  for _, fold := range folds{
    arr := strings.Split(fold, "=")
    axis, value := arr[0], toInt(arr[1])
    switch axis {
    case "x":
      for key := range points {
        if key[0] > value {
          points[[2]int{key[0] - 2 * (key[0] - value), key[1]}] = true
          delete(points, [2]int{key[0], key[1]})
        }
      }
    case "y":
      for key := range points {
        if key[1] > value {
          points[[2]int{key[0], key[1] - 2 * (key[1]-value)}] = true
          delete(points, [2]int{key[0], key[1]})
        }
      }
    }
  }
  
  return len(points)
}

func Part2(input []string) string {
  points, folds := parse(input) 

  for _, fold := range folds{
    arr := strings.Split(fold, "=")
    axis, value := arr[0], toInt(arr[1])

    switch axis {
    case "x":
      for key := range points {
        if key[0] > value {
          points[[2]int{key[0] - 2 * (key[0] - value), key[1]}] = true
          delete(points, [2]int{key[0], key[1]})
        }
      }
    case "y":
      for key := range points {
        if key[1] > value {
          points[[2]int{key[0], key[1] - 2 * (key[1]-value)}] = true
          delete(points, [2]int{key[0], key[1]})
        }
      }
    }
  }

  var maxx int
  var maxy int
  for index := range points {
    if index[0] > maxx {
      maxx = index[0]
    }
    if index[1] > maxy {
      maxy = index[1]
    }
  }

  var sb strings.Builder

  for y := 0; y <= maxy; y++ {
    for x := 0; x <= maxx; x++ {
      point := points[[2]int{x, y}]
      if point {
        sb.WriteRune('▓')
      } else {
        sb.WriteRune(' ')
      }
    }
    sb.WriteString("\n")
  }
   return sb.String()
}

func parse(input []string) (map[[2]int]bool, []string) {
  var modeFold bool
  folds := make([]string, 0)
  points := make(map[[2]int]bool)
  for _, line := range input {
    if line == "" {
      modeFold = true
      continue
    }
    if !modeFold {
      arr := strings.Split(line, ",")
      points[[2]int{toInt(arr[0]), toInt(arr[1])}] = true
    } else {
      folds = append(folds, strings.Fields(line)[2])
    }
  }

  return points, folds
}

func toInt(str string) int {
  num, err := strconv.Atoi(str)

  if err != nil {
    panic(err)
  }

  return num
}
