package samples

import (
    "fmt"
)

// Run ...
func Run() {
    
    text := "this is sample string"
    bytes := []byte(text)
    
    fmt.Print(bytes)
    
    fmt.Println(string(bytes))
}