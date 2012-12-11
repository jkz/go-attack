package ticker

import "time"

var interval = time.Second / 10

func Run(ch chan int) {
    for _ = range time.Tick(time.Second / 10) {
        ch <- 1
    }
}

