#include <stdio.h>

extern void print(const char *s) {
    fputs(s, stdout);
}