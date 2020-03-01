package main
 
import (
    "fmt"
    "math"
)
 
func main(){
    var monsterLife float64
    fmt.Scan(&monsterLife)
    log := math.Floor(math.Log2(monsterLife))
    fmt.Println(int(math.Pow(2, log + 1) - 1))
}
