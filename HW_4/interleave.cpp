#include <iostream>
#include <string>
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
vector<string> vectorInterleave(vector<string> a, vector<string> b) {
    unsigned length = a.size() + b.size();
    vector<string> result;
    int i = 0;
    int j = 0;
    int location = 0;
    while (location < length) {
        if (i < a.size()) {
            result.push_back(a[i]);
            i++;
            location++;
        }

        if (j < b.size()) {
            result.push_back(b[j]);
            j++;
            location++;
        }
    }
    return result;
}

int main() {
    char a[5] = "Dude";
    char b[8] = "1234567";
    std::vector<std::string> c(5, "Mo");
    std::vector<std::string> d(5, "Mo");
    cout << interleave(a,4,b,7) << "\n";
    copy(path.begin(), path.end(), ostream_iterator<string>(cout, " "));
    //std::cout << vectorInterleave(c,d) << "\n";
    cout << interleave(b,7,a,4) << "\n";
    //std::cout << vectorInterleave(d,c) << "\n";
    return 0;
}