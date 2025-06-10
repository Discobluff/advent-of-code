#include "../../utils/parse.h"
#include <stdio.h>

int part1(void){
    int size;
    int *numbers = splitFileToI("2021/01/input.txt",'\n',&size, true);
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
    int size;
    int *numbers = splitFileToI("2021/01/input.txt",'\n',&size, true);
    int res = 0;
    int sum = numbers[0] + numbers[1] + numbers[2];
    for (int i=3;i<size;i++){        
        int sumTemp = sum - numbers[i-3] + numbers[i];
        if (sumTemp > sum){
            res++;
        }
        sum = sumTemp;
    }    
    free(numbers);
    return res;
}

int main(void){
    printf("Part1 : %d\n",part1());
    printf("Part2 : %d\n",part2());
    return 0;
}
