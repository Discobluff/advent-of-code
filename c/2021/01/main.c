#include "../../utils/parse.h"
#include <stdio.h>

int part1(void){
    int size;
    int *numbers = splitFileToI("2021/01/input.txt",'\n',&size);
    int res = 0;
    for (int i=1;i<size;i++){
        if (numbers[i] >= numbers[i-1]){            
            res++;
        }
    }
    free(numbers);
    return res;
}

int part2(void){
    return 0;
}

int main(void){
    printf("Part1 : %d\n",part1());
    printf("Part2 : %d\n",part2());
    return 0;
}
