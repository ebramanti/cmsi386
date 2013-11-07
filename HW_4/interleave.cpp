#include <iostream>
using namespace std;

char* interleave(char a[], char b[]) {
    unsigned length = strlen(a) + strlen(b);
    char* result = new char[length + 1];
    result[length] = '\0';

    int i = 0;
    int j = 0;
    int location = 0;
    while (location < length) {
        if (i < strlen(a)) {
            result[location] = a[i];
            i++;
            location++;
        }

        if (j < strlen(b)) {
            result[location] = b[j];
            j++;
            location++;
        }
    }
    return result;
}

int main() {
    char a[5] = "Dude";
    char b[8] = "1234567";
    cout << interleave(a,b) << "\n";
    return 0;
}