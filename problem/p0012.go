// go version go1.6.2 darwin/amd64
package main

import (
  "fmt"
  "math"
)

func main() {

  var number int = 500
  // 十分に大きな値
  var prime_limit = 10000

  // get primes less than 'number'
  var plist []int = sieveOfEratosthenes(prime_limit)

  var count int = 1;

  for i:=1; ;i+=count {
    if getDivisorCount(i, plist) > number {
      fmt.Println(i)
      break
    }
    count++
  }
}

func getDivisorCount(number int, plist []int) int {
  var count int = 1;
  var flist map[int]int = primeFactorization(number, plist)

  // fmt.Println(number)
  // fmt.Println(flist)

  for _, value := range flist {
    count *= value + 1
  }
  return count
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
