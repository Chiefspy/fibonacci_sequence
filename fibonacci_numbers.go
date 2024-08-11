package main

import "fmt"


func main() {
  max := 100
  a, b := 0 , 1
  for b <= max {
    fmt.println(b)
    a, b = b, a + b
  }

}


