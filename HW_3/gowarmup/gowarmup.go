/*
 *	CMSI 386 HW #3
 *	Edward Bramanti and Zane Kansil
 *
 *	Problem #1
 */

package gowarmup

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
)

/*
 *	A helper method for Change() that performs a Divmod operation.
 */
func DivMod(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}

/*
 * 	A helper method for PowersOfTwo() and Powers() that prints
 * 	results up to maximum value.
 */
func Print(i int) {
	fmt.Println(i)
}

//	Problem a
func Change(amount int) (quarters, nickels, dimes, pennies int) {
	changeLeft := amount
	if amount < 0 || amount%1 != 0 {
		fmt.Println("Non-negative integer expected.")
		panic(amount)
	}
<<<<<<< HEAD
	quarters, changeLeft = Divmod(changeLeft, 25)
	dimes, changeLeft = Divmod(changeLeft, 10)
	nickels, changeLeft = Divmod(changeLeft, 5)
	pennies = changeLeft
=======
	quarters, changeLeft := DivMod(changeLeft, 25)
	dimes, changeLeft := DivMod(changeLeft, 10)
	nickels, changeLeft := DivMod(changeLeft, 5)
	pennies := changeLeft

	result := []int{quarters, nickels, dimes, pennies}

	return result
>>>>>>> 129030e5da3d986b08e17aa355e3b1625d4d9fbb

	return
}

//	Problem b
func RemoveVowels(s string) (t string) {
	return regexp.MustCompile("[aeiouAEIOU]").ReplaceAllString(s, "")
}

//	Problem c
func Scramble(s string) (t string) {
	a := strings.Split(s, "")
	for i := len(a) - 1; i > 0; i-- {
		j := int(math.Floor(rand.Float64() * float64(i)))
		save := a[i]
		a[i] = a[j]
		a[j] = save
	}
	return strings.Join(a, "")
}

//	Problem d
func PowersOfTwo(upTo int, f func(int)) {
	for i := 1; i <= upTo; i *= 2 {
		f(i)
	}
}

// 	Problem e
func Powers(base int, upTo int, f func(int)) {
	if base == 1 {
		f(1)
	} else {
		for i := 1; i <= upTo; i *= base {
			f(i)
		}
	}
}

//	Problem f
func Interleave(a, b []string) (c []string) {
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

//	Problem g
func Stutter(a []string) (c []string) {
	return Interleave(a, a)
}
