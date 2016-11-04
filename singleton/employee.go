package singleton

import (
    "sync"
    "fmt"
)

// Employee ...
type Employee struct {
   ID int
   Name string
}

var instance Employee
var once sync.Once

// GetInstance ...
func GetInstance() Employee {

    once.Do(func() {
        instance = Employee {100, "Ghazni"}
    })
    return instance
}

// Run ...
func Run() {

    fmt.Println(GetInstance())
    fmt.Println(GetInstance()) 
}