package main
import (
  "fmt"
)
func main() {
  
  var a rune = 255                       // rune == int32 => -214748e49 to 214748e48
  fmt.Println(a)

  // intn => -2^(n-1) to 2^(n-1) - 1
  var b int8 = -128                      // -128 to 127
  fmt.Println(b)
  var c int16 = 32767                    // -32,768 to 32,767
  fmt.Println(c)
  var d int32 = -2147483648              // -2,147,483,648 to 2,147,483,647
  fmt.Println(d)
  var e int64 = 9223372036854775807      // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
  fmt.Println(e)
  var f int = 255
  fmt.Println(f)

  // uintn => 0 to 2^(n) - 1
  var g uint8 = 255                     // 0 to 255
  fmt.Println(g)
  var h uint16 = 65535                  // 0 to 65,535
  fmt.Println(h)
  var i uint32 = 4294967295             // 0 to 4,294,967,295
  fmt.Println(i)
  var j uint64 = 18446744073709551615   // 0 to 18,446,744,073,709,551,615 
  fmt.Println(j)
  var k uint = 255
  fmt.Println(k)









}
