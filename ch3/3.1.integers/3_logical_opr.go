package main
import (
  "fmt"
)

func main() {
  fmt.Println("LeftShift")
  var x uint8 = 10
  y := x << 1
  z := x << 5

  fmt.Printf("%d\t%b\n", x, x)
  fmt.Printf("%d\t%b\n", y, y)
  fmt.Printf("%d\t%b\n", z, z)


  fmt.Println("RightShift")
  a := x
  b := a >> 1
  c := a >> 5

  fmt.Printf("%d\t%b\n", a, a)
  fmt.Printf("%d\t%b\n", b, b)
  fmt.Printf("%d\t%b\n", c, c)


  fmt.Println("bitwise OR 1<<1 | 1<<5 ")
  
  fmt.Printf("%d\t%b\n", 1<<1, 1<<1)
  fmt.Printf("%d\t%b\n", 1<<5, 1<<5)
  fmt.Printf("%d\t%b\n", 1<<1|1<<5, 1<<1|1<<5)


  fmt.Println("bitwise bitwise AND 1<<1 & 1<<5 ")
  fmt.Printf("%d\t%b\n", 1<<1&1<<5, 1<<1&1<<5)

  fmt.Println("bitwise bitwise XOR 1<<1 ^ 1<<5 ")
  fmt.Printf("%d\t%b\n", 1<<1^1<<5, 1<<1^1<<5)


  fmt.Println("bitwise bitwise clear 1<<1 &^ 1<<5 ")
  fmt.Printf("%d\t%b\n", 1<<1&^1<<5, 1<<1&^1<<5)


}
