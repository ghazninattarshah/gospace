package samples

import (
    "math/rand"
    "fmt"
    "sort"
)

// Run ...
func Run() {

    arr := make([]int, 100)
    
    for i := 0; i < 100; i++ {
        arr[i] = int(rand.Int31n(100))
        //arr[i] = i + 1
    }
    
    
    sort.Ints(arr)
    fmt.Println(arr)
    
    newArr := make([]int, 10)
    copy(newArr[1:9], arr[:10])
    
    fmt.Println(newArr)
}