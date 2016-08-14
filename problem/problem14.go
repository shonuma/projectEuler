// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
)

func main() {

  // 全探索の終了値
  var number_limit int = 1000000
  m := map[int]int{ 0:0 }

  var step int = 0;
  // 開始値は2
  for i:=2; i < number_limit; i+=1{
    for j:=i; j!=1;  {
      step += 1
      if j % 2 == 0 {
        j = j / 2
      } else {
        j = 3 * j + 1
      }
      if j == 1 {
        m[i] = step
        break
      }
      _, ok := m[j]
      if ok {
        // 途中で探索済みの値が見つかったら終了
        m[i] = step + m[j]
        j = 0
        break
      }
    }
    step = 0
  }
  var max_key int = 0
  var max_val int = 0
  for key,value := range m {
    if value > max_val {
      max_key = key
      max_val = value
    }
  }
  // fmt.Println(m)
  fmt.Println(max_key)
  fmt.Println(max_val)
}
