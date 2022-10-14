package day10

import (
  "sort"
)

func Part1(input []string) int {
  var score int

  for _, line := range input {
    stack := make([]rune, 0)
    for _, char := range line {
      if char == ')' || char == ']' || char == '}' || char == '>' {
        if !(len(stack) > 0 && (char - stack[len(stack)-1] == 1 || char - stack[len(stack)-1] == 2)){
          switch char {
          case ')':
            score += 3
          case ']':
            score += 57
          case '}':
            score += 1197
          case '>':
            score += 25137
          }
        }
        stack = stack[:len(stack)-1]
      } else {
        stack = append(stack, char)
      }
    }

  }

  return score
}

func Part2(input []string) int {
  scores := make([]int, 0)

  discard:
  for _, line := range input {
    stack := make([]rune, 0)
    for _, char := range line {
      if char == ')' || char == ']' || char == '}' || char == '>' {
        if !(len(stack) > 0 && (char - stack[len(stack)-1] == 1 || char - stack[len(stack)-1] == 2)){
          continue discard
        }
        stack = stack[:len(stack)-1]
      } else {
        stack = append(stack, char)
      }
    }
    scores = append(scores, mirrorBrackets(string(stack)))

  }

  sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })

  return scores[(len(scores)/2)]
}

func mirrorBrackets(str string) int {
  var score int

  for i := len(str)-1; i >= 0; i-- {
    score *= 5
    switch str[i] {
    case '(':
      score += 1
    case '[':
      score += 2
    case '{':
      score += 3
    case '<':
      score += 4
    }
  }
  
  return score
}
