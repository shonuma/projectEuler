// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
  "math"
)

func main() {

  // condition: 0<x<y<z, x+y+z = 1000
  var left = 3
  var right = 997
  var flag = 0

  var x,y int
  for ; right>1; right -= 1{
    // x^2 + y^2 < (x+y)^2
    if left < right {
      left += 1
      continue
    }

    // 半分以降は見る必要なし
    for i:=1; i <= int(left/2 + 1); i+=1 {
      x = i
      y = left - i
      if isPythagorean(x, y, right) {
        flag = 1
        break
      }
    }
    if flag == 1 {
      break
    }
    left += 1
  }
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(right)
  fmt.Println(x * y * right)
}

func isPythagorean(x int, y int, z int) bool {
  return int(math.Pow(float64(x), 2.0) + math.Pow(float64(y), 2.0)) == int(math.Pow(float64(z), 2.0))
}
