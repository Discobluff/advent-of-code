#include "../../utils/parse.h"
#include "day.h"
#include <stdio.h>

int min(int *tab, int size){
    int res = tab[0];
    for (int i=1;i<size;i++){
        if (res > tab[i]){
            res = tab[i];
        }
    }
    return res;
}

int max(int *tab, int size){
    int res = tab[0];
    for (int i=1;i<size;i++){
        if (res < tab[i]){
            res = tab[i];
        }
    }
    return res;
}

int abs(int a){
    if (a<0){
        return -a;
    }
    return a;
}

int getFuel1(int *numbers, int size, int score){
    int res = 0;
    for (int i=0;i<size;i++){
        res += abs(numbers[i] - score);
    }
    return res;
}

int getFuel2(int *numbers, int size, int score){
    int res = 0;
    for (int i=0;i<size;i++){
        int a = abs(numbers[i] - score);
        res += a*(a+1)/2;
    }
    return res;
}

int part1(const char* path){
    int size;
    int *numbers = splitFileToI(path, ',', &size, true);
    int mini = min(numbers, size);
    int maxi = max(numbers, size);
    int res = getFuel1(numbers, size, mini);
    for (int i=mini+1; i<=maxi;i++){
        int temp = getFuel1(numbers, size, i);
        if (res > temp){
            res = temp;
        }
    }
    free(numbers);
    return res;
}

int part2(const char* path){    
    int size;
    int *numbers = splitFileToI(path, ',', &size, true);
    int mini = min(numbers, size);
    int maxi = max(numbers, size);
    int res = getFuel2(numbers, size, mini);
    for (int i=mini+1; i<=maxi;i++){
        int temp = getFuel2(numbers, size, i);
        if (res > temp){
            res = temp;
        }
    }
    free(numbers);
    return res;
}
