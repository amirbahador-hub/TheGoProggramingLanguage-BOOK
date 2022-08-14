package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {

    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() { // false if use ctrl + d 
      counts[input.Text()]++

    }

    // Printing The OutPut
    for line, n := range counts {
      if n > 1 {
        fmt.Printf("%d\t%s\n", n, line)
      }
    }

}
