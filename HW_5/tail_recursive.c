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
