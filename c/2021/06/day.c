#include "../../utils/parse.h"
#include "../../utils/grid_lld.h"
#include "day.h"
#include <stdio.h>
#include <assert.h>

int numberFish(int timer, int days){
    if (days == 0){
        return 1;
    }
    if (timer == 0){
        return numberFish(6, days-1) + numberFish(8, days-1);
    }
    return numberFish(timer-1, days-1);
}

int result(const char* path, int days){
    int size;
    int *numbers = splitFileToI(path, ',', &size, true);
    int res = 0;
    for (int i=0;i<size;i++){
        res += numberFish(numbers[i], days);
    }
    free(numbers);
    return res;
}

int part1(const char* path){
    int size;
    int *numbers = splitFileToI(path, ',', &size, true);
    int res = 0;
    for (int i=0;i<size;i++){
        res += numberFish(numbers[i], 80);
    }
    free(numbers);
    return res;
}

long long int part2(const char* path){    
    int size;
    int *numbers = splitFileToI(path, ',', &size, true);
    int totalDay = 256;
    int timerMax = 8;
    int sizes[2] = {totalDay+1, timerMax+1};
    Grid_lld *grid = initGrid_lld(2, sizes, 1);
    for (int day = 1; day<=totalDay; day++){
        for (int timer=0;timer<=timerMax;timer++){
            int pos1[2] = {day, timer};
            if (timer == 0){
                int pos2[2] = {day - 1,6};
                int pos3[2] = {day - 1,8};
                setGrid_lld(grid, pos1, getGrid_lld(grid, pos2) + getGrid_lld(grid, pos3));                
            }
            else{
                int pos2[2] = {day -1, timer - 1};
                setGrid_lld(grid, pos1, getGrid_lld(grid, pos2));                
            }            
        }

    }
    long long int res = 0;
    for (int i=0;i<size;i++){        
        int pos[2] = {totalDay, numbers[i]};
        res += getGrid_lld(grid, pos);
    }
    freeGrid_lld(grid);
    free(numbers);
    return res;
}
