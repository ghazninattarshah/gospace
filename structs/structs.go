package samples

import (
    
    "fmt"
    "strconv"
)

// Employee struct...
type Employee struct {
    ID int
    Name string
}

// Run ...
func Run() {
    
    // emps:= []Employee{}

    // for  i := 0; i < 10; i++ {
    //     em := &Employee{i, "sbros"}
    //     emps = append(emps, *em)
    // }

    emps := createEmployeesSlice()
    fmt.Println(emps)    
    
    emps2 := createEmployeeMap()
    fmt.Println(emps2)
}

// Employee ...
func createEmployeesSlice() []Employee {

    emps:= []Employee{}
    for  i := 0; i < 10; i++ {
        em := &Employee{i, "sbros"}
        emps = append(emps, *em)
    }
    return emps
}

func createEmployeeMap() map[string]Employee {

    emps := make(map[string]Employee)
    
    for i := 0; i < 10; i++ {
        name := "name" + strconv.Itoa(i)
        em := Employee{i, name }
        emps[em.Name] = em
    }
    return emps
}