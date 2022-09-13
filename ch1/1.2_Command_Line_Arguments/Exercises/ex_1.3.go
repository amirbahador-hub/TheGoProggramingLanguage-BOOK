package main

import (
  "fmt"
  "strings"
  "os"
  "time"
)

func main () {

  start := time.Now()

  var s string;
  for i := 1; i<len(os.Args); i++ {
    s += os.Args[i] + " "
  }
  fmt.Printf("%s => for loop %fs \n", s, time.Since(start).Seconds())

  start = time.Now()
  fmt.Printf("%s => Join %fs \n", strings.Join(os.Args[1:], " "), time.Since(start).Seconds())

}
