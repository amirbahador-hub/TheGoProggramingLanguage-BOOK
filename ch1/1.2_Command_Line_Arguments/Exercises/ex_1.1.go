package main

import (
  "fmt"
  "os"
  "strings"
)

func main () {
  fmt.Println(strings.Join(os.Args, " "))
  // OUTPUT with run                       => /tmp/go-build3693683964/b001/exe/ex_1.1 slm hi yw
  // OUTPUT with build then run executable => ./ex_1.1 slm hi yw
}
