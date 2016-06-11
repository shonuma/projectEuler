// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
)

func main() {
  var fib_list = []int{1, 2}
  var sum int = 2
  var limit int = 4000000
  var number int = 0

  for {
    fiblen := len(fib_list)
    number = fib_list[fiblen-1] + fib_list[fiblen-2]
    fib_list = append(fib_list, number)
    if number > limit {
      break
    } else if number % 2 == 0 {
      sum = sum + number
    }
  }

  fmt.Println(sum)
}

