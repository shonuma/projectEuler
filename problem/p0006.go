// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
  "math"
)

func main() {
  var number int = 100

  var square_add = 0;
  var add_square = 0;

  for i:=1;i<=number;i+=1 {
    square_add += int(math.Pow(float64(i), 2.0))
    add_square += i
  }
  add_square = int(math.Pow(float64(add_square), 2.0))

  fmt.Println(square_add)
  fmt.Println(add_square)
  fmt.Println(add_square - square_add)

}
