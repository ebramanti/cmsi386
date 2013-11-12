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
    cout << interleave(a,4,b,7) << "\n";
    cout << interleave(b,7,a,4) << "\n";
    vector <string> c;
    c.push_back("Hello");
    c.push_back("GoodBye");
    vector <string> d;
    d.push_back("Turtles");
    d.push_back("Toolbags");
    vector<string> result_string = vectorInterleave(c,d);
    for(int i = 0; i < result_string.size(); i++){
      cout << result_string[i] << endl;
    }
    return 0;
}