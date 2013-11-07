#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* interleave(char* a, char* b) {
    unsigned length = strlen(a) + strlen(b);
    char* result = malloc(length + 1);
    result[length] = '\0';
    printf("%d\n", length);
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
        printf("%s\n", result);
    }
    return result;
}

int main() {
    char a[4] = "Sup";
    char b[5] = "Dawg";
    printf("Here's the result: %s\n", interleave(a,b));
    return 0;
}