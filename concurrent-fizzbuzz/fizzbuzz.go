package main


import (
    "fmt"
    "strconv"
)


// sentinel value for channels
var EOC int = 0


func responder(in chan int, out chan string) {
    for {
        data := <- in

        // break if you hit sentinel
        if data == EOC {
            out <- string(EOC)
            break
        }

        // switch to send data to out chan
        switch {
            case data % 3 == 0 && data % 5 == 0:
                out <- "fizzbuzz"
            case data % 3 == 0:
                out <- "fizz"
            case data % 5 == 0:
                out <- "buzz"
            default:
                out <- strconv.Itoa(data)
        }
    }
}


func requester(in chan int) {
    // send a seqence of values into in chan
    for i := 1; i < 22; i++ {
        in <- i
    }

    // signal termination of channel with sentinel
    in <- EOC
}


func main() {
    // create in and out channels
    in := make(chan int)
    out := make(chan string)

    // fire off requester and responder
    go requester(in)
    go responder(in, out)

    // read from out channel and display to screen
    for {
        data := <- out

        if data != string(EOC) {
            fmt.Printf("%s ", data)
        } else {
            break
        }
    }
}
