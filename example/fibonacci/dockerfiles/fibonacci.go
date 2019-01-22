package main

import "fmt"
import "flag"
import "os"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func main() {
    fibNum := flag.Int("fibNum", 4, "fibonacci number")
    flag.Parse()
    
    f := fib()
    // Function calls are evaluated left-to-right.
    //fmt.Println(f(), f(), f(), f(), f())
    for a := 0; a < *fibNum; a++ {
        fmt.Println(f())
    }
    os.Exit(0)
}

