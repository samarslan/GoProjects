package main

import (
	"fmt"
  "unsafe"
)

func main() {
  var maxInt uint64 = 1<<64 - 1
  var number *int = nil
  fmt.Printf("Type: %T, Value: %v\n, Hex: %x\n", maxInt, maxInt, maxInt)
  fmt.Println( unsafe.Sizeof( number ) )
}
