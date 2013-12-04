//http://play.golang.org/p/0gyqEAuqhn

package main

import "fmt"

func main() {
    a := []int{1,2,3}
    b := []int{21,7,12,2}
    c := []int{1, 12, -1, 0}
    d := []int{1, 20, 500, 6, -500, 40, -500, 67}
	
    findMinValue(a, 3)
    findMinValue(b, 4)
    findMinValue(c, 4)
    findMinValue(d, 8)
}


func findMinValue(a []int, length int) {
    minValueHelper(a, length, 0, a[0])
}

func minValueHelper(a []int, length int, i int, sofar int) int {
    if (i == length - 1) {
    	if a[i] < sofar {
	    fmt.Println(a[i]) 
 	    return a[i]
	} else {
	    fmt.Println(sofar) 	
	    return sofar
	}
    } else if (a[i] < sofar) {
        sofar = a[i]
    }
    
    return minValueHelper(a, length, i + 1, sofar)
}