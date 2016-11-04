package samples

import (
    
    "fmt"
    "strings"
)

// Employee struct...
type Employee struct {
    ID int
    Name string
}

// Run ...
func Run() {
    
    arr := []int{1,2,3,4,5}
    slice := arr[2:4]
    slice[0] = 50
    
    for index, value := range createEmployees() {
        fmt.Println ("Index =", index, ", ID = ", value.ID, ", Name = ", value.Name)
    }
    
    fmt.Println(strings.Index("this is new world"[5:], " "))
    
    
    fmt.Println()
}

func createEmployees() []Employee {

    emps:= []Employee{}
    for  i := 0; i < 10; i++ {
        em := &Employee{i, "sbros"}
        emps = append(emps, *em)
    }
    
    return emps
}
