package main

import "fmt"            //println stuff
import "os"             //needed to get to Args
import "strconv"        //string conversion methods

func main() {
    n, error := strconv.Atoi(os.Args[1])        //Watch out for methods w multiple returns like: value, error 
    if error != nil {
        fmt.Println("A single integer commandline argument is required")
    } else {
        a, b := 0, 1        //Has multiple assignment.
        for b <= n {
            fmt.Println(b)
            a, b = b, a + b
        }
    }
}