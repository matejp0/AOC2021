package day14

import (
	"strings"
)

func Part1(input []string) uint64 {
  template, rules := parse(input)

  const steps = 10

  for step := 0; step < steps; step++ {
    var sb strings.Builder
    sb.WriteByte(template[0])
    for p := 0; p < len(template)-1; p++ {
      if elem, contains := rules[template[p:p+2]]; contains {
        sb.WriteString(elem)
      }
      sb.WriteByte(template[p+1])
    }
    template = sb.String();
  }

  c := counts(template)

  var leastCommon = uint64(len(template))
  var mostCommon uint64
  for _, value := range c {
    if value > mostCommon {
      mostCommon = value
    }
    if value < leastCommon {
      leastCommon = value
    }
  }

  return mostCommon-leastCommon

}

func Part2(input []string) uint64 {
  return process(input, 40)
}

func process(input []string, steps int) uint64 {
  template, rules := parse(input)
  pairs := make(map[string]uint64)

  for p := 0; p < len(template) - 1; p++ {
    pairs[template[p:p+2]]++
  }


  temp := make(map[string]uint64)
  for k, v := range pairs {
    temp[k] = v
  }
  for step := 0; step < steps; step++ {

    for key, value := range pairs {
      if elem, contains := rules[key]; contains && value != 0 {
        temp[key] -= value
        temp[string(key[0]) + elem] += value
        temp[elem + string(key[1])] += value
      }
    }

    for k, v := range temp {
      pairs[k] = v
    }
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


func counts(str string) map[rune]uint64 {
  c := make(map[rune]uint64)

  for _, char := range str {
    c[char]++
  }

  return c
}

func parse(lines []string) (string, map[string]string){
  var template string
  rules := make(map[string]string, 0)
  for i, line := range lines {
    if i == 0 {
      template = line
      continue
    }
    if line == "" { continue }

    arr := strings.Split(line, " -> ")
    rules[arr[0]] = arr[1]

  }

  return template, rules
}

