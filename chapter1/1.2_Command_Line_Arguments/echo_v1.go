package main

import (
  "fmt"
  "os"
)

func main () {

  var s string;
  for i := 1; i<len(os.Args); i++ {
    s += os.Args[i] + " "
  }
  fmt.Println(s)

}
