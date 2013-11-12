#include <stdio.h>
#include <string.h>
#include <math.h>

void setup_a() {
    static int magic_number = 8;
}

void setup_b() {
    static int magic_number = 42;
}

void setup_c() {
    static int magic_number = 777;
}

int main() {
    int r = rand() % 2;
    if (r == 0) {
        setup_a();
    } else if (r == 1) {
        setup_b();
    } else { 
        setup_c();
    }
    printf('%d\n', magic_number);
}



