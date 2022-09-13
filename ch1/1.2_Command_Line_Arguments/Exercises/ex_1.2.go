package main

import (
  "fmt"
  "os"
)

func main () {
  
  for index, arg := range os.Args {
    fmt.Println(index, arg)
  }
}
