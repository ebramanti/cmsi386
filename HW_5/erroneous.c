#include <stdio.h>

void foo() {
    int i;
    printf("%d\n", i++); //added new-line for cleanness
  }
int main() {
    int j;
    for (j = 1; j <= 10; j++) foo();
}