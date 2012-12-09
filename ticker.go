package main

import "fmt"
import "time"
var interval = time.Second / 10

func main() {
    c := time.Tick(time.Second / 10)
    for {
        select {
        case now := <- c:
            fmt.Printf("TICK", now)
        default:
            fmt.Printf(".")
        }
    }
}

