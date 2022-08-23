package main

import (
  "fmt"
  "os"
  "strconv"
  "github.com/amirbahador-hub/TheGoProggramingLanguage-BOOK-/tree/master/chapter2/2.6.Packages_and_Files/temconv"
  //"gopl.io/ch2/tempconv"
)

func main() {
  for _, arg := range os.Args[1:] {
    t, err := strconv.ParseFloat(arg. 64)
    if err != nil {
      fmt.Fprintf(os.Stderr, "cf: %v\n", err)
      os.Exit(1)
    }
    f := tempconv.Fahrenheit(t)
    c := tempconv.Celsius(t)
    fmt.Printf("%s = %s, %s = %s\n",
      f, tempconv.FToC(f), c, tempconv.CToF(c))
  }
}
