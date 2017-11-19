package main

import "fmt"
import "time"

func main() {
ticker := time.NewTicker(5 * time.Second)
quit := make(chan struct{})
    for {
       select {
        case <- ticker.C:
            tick()
        case <- quit:
            ticker.Stop()
            return
        }
    }
}

//TODO: AT - Add calculator function that adds to number
func tick() {
	fmt.Println("Hello World")
}