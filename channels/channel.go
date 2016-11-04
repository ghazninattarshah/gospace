package samples

import (
    "fmt"
)

// Run ...
func Run() {
 
  fmt.Println("channel sample...")

  cht := make(chan string)

  go func ()  {
      cht <- "test"
      cht <- " message"
  }()
  
  for i := 0; i < 2; i++ {
      select {
          default:
          fmt.Println("received", <-cht)
      }
  }
  
  //fmt.Println(<-cht);
  
}