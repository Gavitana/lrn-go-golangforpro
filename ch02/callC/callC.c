#include <stdio.h>
#include "callC.h"

int fibonacci(int num) {
    int a = 0;
    int b = 1;
    int d = 0;
    int c=0;
    for (c=1; c<=num; c = c + 1){
        d = a + b;
        a = b;
        b = d;
    }
    return d;


}
