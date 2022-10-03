package day4

import (
  "strings"
  "strconv"
  //"fmt"
)

const BOARD_SIZE = 5

func Part1(input []string) int {
  numbers := parseNumbers(input[0])
  boards := parseBoards(input)

  for _, number := range numbers {
    for _, board := range boards {
      board.mark(number)
      if(board.containsRowOrColumn()){
        return board.sumUnmarked() * number
      }
    }
  }
  return 0
}

func (board *Board) mark(number int) {
  for _, row := range board.values {
    for _, num := range row {
      if num.value == number {
        num.marked = true
      }
    }
  }

}

func Part2(input []string) int {
  numbers := parseNumbers(input[0])
  boards := parseBoards(input)
  
  wonboards := make([]Board, 0)
  for _, number := range numbers {
    for _, board := range boards {
      board.mark(number)

      if board.containsRowOrColumn() && !board.containedIn(wonboards){
        wonboards = append(wonboards, board)
        if (len(wonboards) == len(boards)) {
          return number*board.sumUnmarked()
        }
      }
    }
  }

  return 0
}
  
type Number struct {
  value int
  marked bool
}

type Board struct {
  values [BOARD_SIZE][BOARD_SIZE]*Number
  number int
}

func (board *Board) containedIn(boards []Board) bool {
  for _, v := range boards {
    if *board == v {
      return true
    }
  }
  return false
}

func (board *Board) containsRowOrColumn() bool {
  var vertical [BOARD_SIZE]int
  for _, row := range board.values {
    var horizontal int
    for i, n := range row {
      if n.marked {
        horizontal++
        vertical[i]++
      }
    }
    if(horizontal == BOARD_SIZE) { return true }
    horizontal = 0
  }
  for _, v := range vertical {
    if v == BOARD_SIZE {
      return true
    }
  }
  return false
}

func (board *Board) sumUnmarked() int {
  var unmarked int
  for _, row := range board.values {
    for _, n := range row {
      if !n.marked {
        unmarked += n.value
      }
    }
  }
  return unmarked
}

func parseNumbers(line string) []int {
  var numbers []int
  for _, v := range strings.Split(line, ",") {
    number, _ := strconv.Atoi(v)
    numbers = append(numbers, number)
  }
  return numbers
}

func parseBoards(input []string) []Board {
  boards := make([]Board, 0)
  var board Board = Board{}
  var rownum int
  for line := 2; line < len(input); line++ {
    if input[line] != "" {
      var row [BOARD_SIZE]*Number
      for i, v := range strings.Fields(input[line]) {
        num, err := strconv.Atoi(v)
        if err != nil { panic(err) }
        row[i] = &Number{value: num, marked: false}
      }
      board.values[rownum] = row
      rownum++

    } else {
      board.number = len(boards)+1
      boards = append(boards, board)
      board = Board{}
      rownum = 0
    }
  }

  board.number = len(boards)+1
  boards = append(boards, board)
  
  return boards
}
