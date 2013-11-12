#include <iostream>
#include <vector>
using namespace std;

//  Problem #9
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

//  Problem #10
char* vectorInterleave(char a[], char b[]) {
    unsigned length = strlen(a) + strlen(b);
    vector<char> result;
    int i = 0;
    int j = 0;
    int location = 0;
    while (location < length) {
        if (i < strlen(a)) {
            result.push_back(a[i]);
            i++;
            location++;
        }

        if (j < strlen(b)) {
            result.push_back(b[j]);
            j++;
            location++;
        }
    }
    /* A vector guarantees that its elements occupy contiguous memory, 
     * so data is at address of the first element.
     */
    return &result[0];
}

int main() {
    char a[5] = "Dude";
    char b[8] = "1234567";
    cout << interleave(a,4,b,7) << "\n";
    cout << vectorInterleave(a,b) << "\n";
    cout << interleave(b,7,a,4) << "\n";
    cout << vectorInterleave(b,a) << "\n";
    return 0;
}