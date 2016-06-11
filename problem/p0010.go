// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
  "math"
)

func main() {
  var limit int = 2000000
  var plist []int = sieveOfEratosthenesByValue(limit)
  var result int = 0

  for i:=0; i< len(plist); i++ {
    result += plist[i]
  }
  // fmt.Println(plist)
  fmt.Println(result)
}

func sieveOfEratosthenesByValue(limit int) []int{
  var plist = []int{2}
  for i:=3; ; i+=1 {
    if i % 2 == 0 {
      continue
    }
    var plen int = len(plist)
    var flag int = 0
    for j:=0; j<plen; j+=1 {
      if i % plist[j] == 0 { 
        flag = 1
        break
      } else if plist[j] > int(math.Sqrt(float64(i))) {
        flag = 0
        break
      }
    }
    if i > limit {
      break
    } else if flag == 0 { 
      plist = append(plist, i)
    }
    flag = 0
  }
  return plist
}
