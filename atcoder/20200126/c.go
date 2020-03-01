package main

import (
    "fmt"
    "sort"
)


func main(){
    var monsterNum int
    fmt.Scan(&monsterNum)
    
    monsterLives := make([]int, monsterNum)
    var hissatsuNum int
    fmt.Scan(&hissatsuNum)
    for i := 0; i < monsterNum; i++ {
        fmt.Scan(&monsterLives[i])
    }
    sort.Sort(sort.Reverse(sort.IntSlice(monsterLives)))

    sum := 0
    for i := hissatsuNum; i < monsterNum; i++ {
        sum = sum + monsterLives[i]
    }
    fmt.Println(sum)
}

