/*
Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  isFirst := true
  cur, prev := 0, 1

  return func() int {
    if isFirst {
      isFirst = false
      return 0
    }

    cur, prev = cur + prev, cur
    return cur
  }
}

func main() {
  f := fibonacci()

  for i := 0; i < 20; i++ {
    fmt.Println(f())
  }
}
