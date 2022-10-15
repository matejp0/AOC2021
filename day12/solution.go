package day12

import (
	"strings"
)

const START = "start"
const END = "end"

func Part1(input []string) int {
  return dfs(parse(input), START, make(map[string]bool))
}

func Part2(input []string) int {
  return dfs2(parse(input), START, make(map[string]uint8))
}

func dfs(arr map[string][]string, current string, visited map[string]bool) int {
  var count int

  if current == END {
    return 1
  }

  if visited[current] {
    return 0
  }

  if isSmall(current) {
    visited[current] = true
  }

  for _, val := range arr[current] {
    if val != START {
      count += dfs(arr, val, visited)
    }
  }
  
  if isSmall(current) {
    delete(visited, current)
  }

  return count
}

func dfs2(arr map[string][]string, current string, visited map[string]uint8) int {
  var count int

  if current == END {
    return 1
  }

  if isSmall(current) {
    var twice int
    for _, v := range visited {
      if v > 1 {
        twice++
      }
    }
    if twice == 0 && visited[current] <= 1 {
      visited[current]++
    } else if twice == 1 && visited[current] == 0 {
      visited[current]++
    } else {
      return 0
    }
  }

  for _, val := range arr[current] {
    if val != START {
      count += dfs2(arr, val, visited)
    }
  }
  
  if isSmall(current) {
    visited[current]--
  }

  return count
}


func isSmall(str string) bool {
  return strings.ToLower(str) == str
}

func parse(input []string) map[string][]string {
  result := make(map[string][]string, 0)
  for _, line := range input {
    arr := strings.Split(line, "-")
    
    result[arr[0]] = append(result[arr[0]], arr[1])
    result[arr[1]] = append(result[arr[1]], arr[0])
  }
  return result
}

// https://go.dev/blog/maps - handy reference
