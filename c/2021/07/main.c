#include "day.h"
#include <stdio.h>
#include <time.h>

int main(void){
    
    clock_t start, end;
    double cpu_time_used;

    start = clock();
    printf("Part1 : %d\n", part1("07/input.txt"));
    end = clock();
    cpu_time_used = ((double)(end - start)) / CLOCKS_PER_SEC;
    printf("Part1 execution time: %f seconds\n", cpu_time_used);

    start = clock();
    printf("Part2 : %d\n", part2("07/input.txt"));
    end = clock();
    cpu_time_used = ((double)(end - start)) / CLOCKS_PER_SEC;
    printf("Part2 execution time: %f seconds\n", cpu_time_used);    
    return 0;
}
