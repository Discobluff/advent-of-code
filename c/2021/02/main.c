#include "../../utils/parse.h"
#include <stdio.h>

int part1(void){
    int size;
    char **instructions = splitFile("2021/02/input.txt",'\n',&size, true);
    int abs = 0;
    int depth = 0;
    for (int i=0;i<size;i++){
        int command;
        char c[10];
        sscanf(instructions[i], "%s %d", c, &command);
        if (instructions[i][0] == 'u'){
            depth -= command;
        }
        if (instructions[i][0] == 'f'){
            abs += command;
        }
        if (instructions[i][0] == 'd'){
            depth += command;
        }
    }
    free(instructions);
    return abs*depth;
}

int part2(void){
    int size;
    char **instructions = splitFile("2021/02/input.txt",'\n',&size, true);
    int abs = 0;
    int depth = 0;
    int aim = 0;
    for (int i=0;i<size;i++){
        int command;
        char c[10];
        sscanf(instructions[i], "%s %d", c, &command);
        if (instructions[i][0] == 'u'){
            aim -= command;
        }
        if (instructions[i][0] == 'f'){
            abs += command;
            depth += aim*command;
        }
        if (instructions[i][0] == 'd'){
            aim += command;
        }
    }
    free(instructions);
    return abs*depth;
}

int main(void){
    printf("Part1 : %d\n",part1());
    printf("Part2 : %d\n",part2());
    return 0;
}
