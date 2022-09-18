package main
import  (
  "fmt"
)

func main(){

  var a rune = 2147483647                // rune == int32 => -214748e49 to 214748e48
  fmt.Println(a, a+1, a*a) 

  // intn => -2^(n-1) to 2^(n-1) - 1
  var b int8 = -128                      // -128 to 127
  fmt.Println(b, b-1, b*b) 
  var c int16 = 32767                    // -32,768 to 32,767
  fmt.Println(c, c+1, c*c) 
  var d int32 = -2147483648              // -2,147,483,648 to 2,147,483,647
  fmt.Println(d, d-1, d*d) 
  var e int64 = 9223372036854775807      // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
  fmt.Println(e, e+1, e*e) 
  var f int = -2147483648                // int == int32 same size but not same type
  fmt.Println(f, f-1, f*f) 

  // uintn => 0 to 2^(n) - 1
  var g uint8 = 255                     // 0 to 255
  fmt.Println(g, g+1, g*g) 
  var h uint16 = 65535                  // 0 to 65,535
  fmt.Println(h, h+1, h*h) 
  var i uint32 = 4294967295             // 0 to 4,294,967,295
  fmt.Println(i, i+1, i*i) 
  var j uint64 = 18446744073709551615   // 0 to 18,446,744,073,709,551,615 
  fmt.Println(j, j+1, j*j) 
  var k uint = 255                      // unit == unit32 same size but not same type
  fmt.Println(k, k+1, k*k) 


}
