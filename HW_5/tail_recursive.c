#include <stdio.h>

void findMinValue(int * anArray, int length) {
	minValueHelper(anArray, length, 0, anArray[0]);
}

int minValueHelper (int * anArray, int length, int i, int sofar) {
    if (i == length - 1) {
        return  (a[i] < sofar) ? a[i] : sofar;	
    }
    else if (a[i] < sofar) {
        sofar = a[i]
    }
    
    return findMinValue(a, i + 1, sofar);
}



int main() {
	int a[3] = {1, 2, 3};
	int b[4] = {21,7,12,2};
	int c[4] = {1, 12, -1, 0};
	int d[8] = {1, 20, 500, 6, -500, 40, -500, 67};
	findMinValue(a, 3);
	findMinValue(b, 4);
	findMinValue(c, 4);
	findMinValue(d, 8);
}