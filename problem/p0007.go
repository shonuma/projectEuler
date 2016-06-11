// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
)

func main() {
  var count int = 10001

  fmt.Println(sieveOfEratosthenesByIndex(count))

}

func sieveOfEratosthenesByIndex(count int) int{
  var plist = []int{2}
  var result int = 0
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
      } 
    }
    if flag == 0 { 
      plist = append(plist, i)
      if len(plist) == 10001 {
        result = i
        break
      }
    }
    flag = 0
  }
  return result
}
