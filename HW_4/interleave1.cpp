char* interleave(char a[], int aLength, char b[], int bLength) {
    unsigned length = aLength + bLength;
    char* result = new char[length + 1];
    result[length] = '\0';
    int i = 0;
    int j = 0;
    int location = 0;
    while (location < length) {
        if (i < aLength) {
            result[location] = a[i];
            i++;
            location++;
        }

        if (j < bLength) {
            result[location] = b[j];
            j++;
            location++;
        }
    }
    return result;
}