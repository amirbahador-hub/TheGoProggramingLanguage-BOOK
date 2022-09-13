package main

import (
  "fmt"
)

func main()  {

  fmt.Println("for loop")
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }

  fmt.Println("while loop")
  i := 0 // we defind i in line 10 BUT not in this scope so we should defind it here
  for i < 5 {
    fmt.Println(i)
    i++
  }

  fmt.Println("infinite loop")
  i = 0 // var i is defind in line 15 so we cannot use :=
  for {
    fmt.Println(i)
    if i == 4 {
      break
    }
    i++
  }

  fmt.Println("range loop")
  my_text := "salam"
  // range will work on slice and string
  for index, _ := range my_text { // range loop
    fmt.Println(index) // we ignored ascci word by using underscore because we did not use it 
  }

}
