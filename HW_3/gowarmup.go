package gowarmup

import (
	"container/list"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
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
	if amount < 0 || amount%1 != 0 {
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
	return strings.Join(a, "")
}

func print(i int) {
	fmt.Println(i)
}

func powersOfTwo(upTo int, f func(int)) {
	for i := 1; i <= max; i *= 2 {
		f(i)
	}
}

func powers(base int, upTo int, f func(int)) {
	if base == 1 {
		f(1)
	} else {
		for i := 1; i <= upTo; i *= base {
			f(i)
		}
	}
}

func interleave(a, b []string) (c []string) {
	var result []string
	i := 0
	j := 0

	for len(result) < (len(a) + len(b)) {
		if i < len(a) {
			result = append(result, a[i])
			i++
		}
		if j < len(b) {
			result = append(result, b[j])
			j++
		}
	}
	return result
}

func stutter(a []string) (c []string) {
	return interleave(a, a)
}
