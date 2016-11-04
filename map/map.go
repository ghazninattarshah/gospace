package samples

import (
    "fmt"
)

// Run ...
func Run() {

    props := make(map[string] string)
    
    props["name"] = "ghazni"
    props["newName"] = "test"
    value, exist := props["name"]
    
    fmt.Println("value = ", value, ",exist = ", exist)
    
    fmt.Println(props)
    fmt.Println(len(props))
    
    delete(props, "newName")
    
    fmt.Println(props)
}