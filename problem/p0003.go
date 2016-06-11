// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
  "math"
)

func main() {
  var number int = 600851475143

  var number_sqrt_max int = int(math.Sqrt(float64(number)) + 1)

  plist := sieveOfEratosthenes(number_sqrt_max)

  for i:=0; i<len(plist); i+=1 {
    if number % plist[i] == 0 {
      fmt.Println(plist[i])
      //fmt.Println(number / plist[i])
      //break
    }
  }

  //fmt.Println(number_sqrt_max)
}

func sieveOfEratosthenes(n int) []int{
  var plist = []int{2}
  for i:=3; i<n; i+=1 {
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
    if flag == 0 { 
      plist = append(plist, i)
    }
    flag = 0
  }
  return plist
}

