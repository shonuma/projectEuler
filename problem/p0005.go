// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
  "math"
)

func main() {

  var number int = 20
  var resultflist map[int]int = map[int]int{} 

  // get primes less than 'number'
  var plist []int = sieveOfEratosthenes(number)

  for i:=1; i<=number; i+=1 {
    // compare current_factorization with result_factorization
    // update result_factorization
    numberflist := primeFactorization(i ,plist)
    for key, value := range numberflist {
      if resultflist[key] < value {
        resultflist[key] = value
        fmt.Println(key)
        fmt.Println(resultflist)
      }
    }
  }

  var result int = 1
  for key,value := range resultflist {
    result *= int(math.Pow(float64(key),float64(value)));
  }
  fmt.Println(result)
}

func primeFactorization(n int, plist []int) map[int]int{
  var factlist map[int]int = map[int]int{}
  for i:=0; i<len(plist); i+=1 {
    prime := plist[i]

    if n % prime == 0 {
      factlist[prime] += 1
      n = n / prime;
      // 同じ値をもう1回見る
      i =- 1;
    }
  }
  return factlist
}

func sieveOfEratosthenes(n int) []int {
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
