package day14

import (
	"strings"
)

func Part1(input []string) uint64 {
  return process(input, 10)
}

func Part2(input []string) uint64 {
  return process(input, 40)
}

func process(input []string, steps int) uint64 {
  pairs, rules := parse(input)

  temp := deepCopy(pairs)

  for step := 0; step < steps; step++ {
    for key, value := range pairs {
      if elem, contains := rules[key]; contains && value != 0 {
        temp[key] -= value
        temp[string(key[0]) + elem] += value
        temp[elem + string(key[1])] += value
      }
    }

    pairs = deepCopy(temp)
  }

  commonLetters := make(map[byte]uint64)
  for key, value := range pairs {
    commonLetters[key[0]] += value 
    commonLetters[key[1]] += value 
  }
  
  var mostCommon uint64
  var leastCommon uint64 = ^uint64(0)

  for _, value := range commonLetters {
    if value > mostCommon { mostCommon = value }
    if value < leastCommon { leastCommon = value }
  }

  return (uint64(mostCommon)-uint64(leastCommon))/2
}


func deepCopy(orig map[string]uint64) map[string]uint64 {
  dest := make(map[string]uint64)
  for k, v := range orig {
    dest[k] = v
  }
  return dest
}

func parse(lines []string) (map[string]uint64, map[string]string){
  rules := make(map[string]string, 0)
  pairs := make(map[string]uint64)
  
  for i, line := range lines {
    if i == 0 {
      for p := 0; p < len(line) - 1; p++ {
        pairs[line[p:p+2]]++
      }
      continue
    }
    if line == "" { continue }

    arr := strings.Split(line, " -> ")
    rules[arr[0]] = arr[1]

  }

  return pairs, rules
}

