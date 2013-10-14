package gowarmup

import (
    "fmt"
    "strings"
    "strconv"
    "regexp"
    "container/list"
    "math"
)

func main() {
    //mini test here
    change(20)
    //..
}

func divmod(x, y int) (quotient, remainder int) {
    quotient = x / y
    remainder = x % y
    return
}

func change(amount int) (change []int) {
    changeLeft := amount
    if (amount < 0 || amount % 1 != 0) {
        //
    }
    quarters, changeLeft := divmod(changeLeft, 25)
    dimes, changeLeft := divmod(changeLeft, 10)
    nickels, changeLeft := divmod(changeLeft, 5)
    pennies := changeLeft

    result := []int{quarters, nickels, dimes, pennies}

    return result

}

func removeVowels(s string) (t string) {
    //
}

func scramble(s string) (t string) {
    //
}

func powersOfTwo(upTo int)  {
    //
}

func powers(upTo, base int) (result int) {
    //
}

func interleave(a, b *List) (c List) {
    //
}

func stutter(a, b *List) (c List) {
    //
}









