package main

import (
  test "golang/test"
)

import (
	"fmt"
)

func main() {
  option := 9
  // menu
  fmt.Println("Run demo: ")
  fmt.Println("----------------- ")
  fmt.Println("0 = quit ")        
  fmt.Println("1 = Hello World")
  fmt.Println("2 = For loop")  
  fmt.Println("----------------- ")
  for ;option>0; {
    fmt.Print(">>")
    fmt.Scanf("%d",&option)
    if option == 0 {
       break 
    } else if option == 1 {
      test.Hello()
    } else if option == 2 {
      test.For_loop()
    } else {
      fmt.Println("error")
      break
    }
  }
}
