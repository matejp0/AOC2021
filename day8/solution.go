package day8

import (
  //"fmt"
  "strings"
  "strconv"
  "sort"
)

func Part1(input []string) int {
  var count int
  for _, line := range input {
    output := strings.Split(line, " | ")[1]
    for _, v := range strings.Fields(output) {
      if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
        count++
      }
    }
  } 
  return count
}

func Part2(input []string) int {
  var sum int
  for _, line := range input {
    ten := strings.Fields(strings.Split(line, " | ")[0])
    output := strings.Fields(strings.Split(line, " | ")[1])
    
    var nums [11]string
    sort.Slice(ten, func(i, j int) bool { return len(ten[i]) < len(ten[j]) }) // sort from the shortest
    for _, v := range ten {
      switch len(v) {
      case 2:
        nums[1] = v
      case 3:
        nums[7] = v
      case 4:
        nums[4] = v

        // Create another char by basically subtracting 1 from 4 (in characters)
        var l string
        for _, char := range v {
          var in1 bool
          for _, char1 := range nums[1] {
            if char == char1 {
              in1 = true
              break
            } 
          }
          if !in1 {
            l += string(char)
          }
        }
        nums[10] = l // helpful char in determining the difference between 2&5 and 9&0
      case 5:
        if containsAll(v, nums[1]) {
          nums[3] = v
        } else if containsAll(v, nums[10]) {
          nums[5] = v
        } else {
          nums[2] = v
        }
      case 6:
        if containsAll(v, nums[1]) {
          if containsAll(v, nums[10]) {
            nums[9] = v
          } else {
            nums[0] = v
          }
        } else {
          nums[6] = v
        }
      case 7:
        nums[8] = v
      }
    }

    var number string
    for _, v := range output {
      num := find(v, nums)
      number += num
    }
    
    numberInt, err := strconv.Atoi(number)
    if err != nil { panic(err) }

    sum += numberInt
  }
  return sum
}


func find(s string, nums [11]string) string {
  for i, num := range nums {
    var count int
    for _, char := range num {
      if strings.ContainsRune(s, char) {
        count++ 
      }
    }
    if len(s) == count && count == len(num) {
      return strconv.Itoa(i)
    }
  }
  return ""
}

func containsAll(num string, subNum string) bool {
  var count int
  for _, char := range num {
    for _, subChar := range subNum {
      if char == subChar {
        count++
      }
    }
  }

  if count == len(subNum) {
    return true
  }
  return false
}
