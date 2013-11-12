#include <stdio.h>

int f(int x) {
    /*static*/ int a = 99;
    if (x == 0) {
        a = 100;
        return f(1);
    }
    else if (x == 1) {
        return a;
    }
    return -1;
}

int main() {
    printf("%d\n", f(0));
    return 0;
}