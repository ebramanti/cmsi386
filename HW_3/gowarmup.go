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
    fmt.Println(change(20))
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
    return regexp.MustCompile("[aeiouAEIOU]").ReplaceAllString(s, "")
}

func scramble(s string) (t string) {
    a := strings.Split(s, "")
    for i := len(a) - 1; i > 0; i-- {
        j := int(math.Floor(rand.Float64() * float64(i)))
        save := a[i]
        a[i] = a[j]
        a[j] = save
    }
    return strings.Join(a,"")
}

func powersOfTwo(upTo int, func(int) string) {
    //http://play.golang.org/p/IyCHPck7ri
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









