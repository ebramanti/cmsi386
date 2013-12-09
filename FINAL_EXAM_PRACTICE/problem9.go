//http://play.golang.org/p/hydj3ZH4IF

package main

import "fmt"

type Sequence []interface{}

func main() {
    fmt.Println(withoutEvens(Sequence{0,1,2,3,4,5,6}, 7))
    fmt.Println(withoutEvens(Sequence{4,6,true,nil,2.3,"xyz"}, 6))
}

func withoutEvens(a Sequence, length int)(Sequence) {
    result := make(Sequence, 0)
    return withoutEvensHelper(a, length, 1, result)
}

func withoutEvensHelper(a Sequence, length int, i int, result Sequence) (Sequence) {
    if (i >= length) {
        return result
    } else {
        result = append(result, a[i])
    }
    return withoutEvensHelper(a, length, i + 2, result)
}