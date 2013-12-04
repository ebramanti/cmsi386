#include <stdio.h>

int minValueHelper (int * anArray, int length, int i, int sofar) {
    if (i == length - 1) {
        return  (anArray[i] < sofar) ? anArray[i] : sofar;  
    }
    else if (anArray[i] < sofar) {
        sofar = anArray[i];
    }
    
    return minValueHelper(anArray, length, i + 1, sofar);
}

int findMinValue(int * anArray, int length) {
    return minValueHelper(anArray, length, 0, anArray[0]);
}

int main() {
    int a[3] = {1, 2, 3};
    int b[4] = {21,7,12,2};
    int c[4] = {1, 12, -1, 0};
    int d[8] = {1, 20, 500, 6, -500, 40, -500, 67};
    printf("%d\n", findMinValue(a, 3));
    printf("%d\n", findMinValue(b, 4));
    printf("%d\n", findMinValue(c, 4));
    printf("%d\n", findMinValue(d, 8));
}