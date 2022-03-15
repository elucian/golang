
package demo

import (
    "fmt"
    "time"
)

func Infinite_loop() {
  for {
    fmt.Println("Press Ctrl+C to stop:")
    time.Sleep(3 * time.Second)
  }
}