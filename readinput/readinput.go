package readinput

import (
  "bufio"
  "fmt"
  "os"
  "log"
)

func Read(number int) []string {
  file, err := os.Open(fmt.Sprintf("inputs/%d", number))
  if err != nil { log.Fatal(err) }
  defer file.Close()

  lines := make([]string, 0)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil { log.Fatal(err) }

  return lines
}
