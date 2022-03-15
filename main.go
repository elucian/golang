package main

import (
  demo "golang/demo"
)

import (
	"fmt"
)

func print_menu() {
  fmt.Println("----------------- ")
  fmt.Println("0 = quit ")        
  fmt.Println("1 = Hello World")
  fmt.Println("2 = For loop")  
  fmt.Println("3 = Random grades")    
  fmt.Println("4 = Infinite loop")    
  fmt.Println("----------------- ")  
}

func main() {
  option := 9
  fmt.Println("Run demo: ")
  print_menu() 
  for ;option>0; {
    fmt.Print(">>")
    fmt.Scanf("%d",&option)
    if option == 0 {
       break 
    } else if option == 1 {
      demo.Hello()
    } else if option == 2 {
      demo.For_loop()      
    } else if option == 3 {
      demo.Random_grades()
    } else if option == 4 {
      demo.Infinite_loop()
    } else {
      fmt.Println("invalid option")
      print_menu() 
    }
  }
}
