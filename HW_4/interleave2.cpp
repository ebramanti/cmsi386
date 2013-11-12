#include <string>
#include <vector>
using namespace std;

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