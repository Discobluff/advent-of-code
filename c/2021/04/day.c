#include "../../utils/parse.h"
#include "day.h"
#include <stdio.h>
#include <string.h>

#define sizeGrid 5

int *parseLineGrid(char *line){
    int *res = malloc(sizeof(int)*sizeGrid);
    int index = 0;
    for (int i=0;i<5;i++){
        int number = 0;        
        if (line[index] == ' '){            
            number = line[index+1] - '0';            
        }
        else{
            number = (line[index] - '0')*10 + line[index+1] - '0';
        }
        index+=3;
        res[i] = number;        
    }
    return res;
}

bool checkGrid(bool **grid){
    for (int line = 0;line<sizeGrid;line++){
        bool res = true;
        for (int column = 0;column<sizeGrid;column++){
            if (!grid[line][column]){
                res  =false;
            }
        }
        if (res){
            return true;
        }
    }
    for (int column = 0;column<sizeGrid;column++){
        bool res = true;
        for (int line = 0;line<sizeGrid;line++){
            if (!grid[line][column]){
                res  =false;
            }
        }
        if (res){
            return true;
        }
    }
    return false;
}

int ***parseGrids(char **lines, int size, int *gridSize){
    int countGrid = 1;
    int indexLine = 2;
    for (int i=2;i<size;i++){
        if (strlen(lines[i]) == 0){
            countGrid++;
        }
    }
    int ***grids = malloc(sizeof(int**)*countGrid);
    int indexGrid = 0;
    while (indexGrid < countGrid){
        int **grid = malloc(sizeof(int*)*sizeGrid);
        for (int i=0;i<sizeGrid;i++){
            grid[i] = parseLineGrid(lines[indexLine]);                        
            indexLine++;
        }
        indexLine++;
        grids[indexGrid] = grid;
        indexGrid++;
    }
    *gridSize = countGrid;
    return grids;
}

bool ***initSeen(int countGrid){
    bool*** seen = malloc(countGrid * sizeof(bool**));
    for (int i=0;i<countGrid;i++){
        seen[i] = malloc(sizeof(bool*)*sizeGrid);
        for (int j=0;j<sizeGrid;j++){
            seen[i][j] = calloc(sizeGrid,sizeof(bool));
        }
    }
    return seen;
}

int part1(const char* path){
    int size;
    char **lines = splitFile(path, '\n',&size,true);
    int countNumbers;    
    char **numbersString = splitString(lines[0],',',&countNumbers,true);
    int *numbers = atoiArray(numbersString,countNumbers);
    int countGrid;
    int ***grids = parseGrids(lines, size, &countGrid);
    bool ***seen = initSeen(countGrid);
    int res = 0;
    for (int i=0;i<countNumbers;i++){
        for (int j=0;j<countGrid;j++){
            for (int k=0;k<sizeGrid;k++){
                for (int l=0;l<sizeGrid;l++){
                    if (grids[j][k][l] == numbers[i]){
                        seen[j][k][l] = true;
                    }
                }
            }
            if (checkGrid(seen[j])){                
                for (int k=0;k<sizeGrid;k++){
                    for (int l=0;l<sizeGrid;l++){
                        if (!seen[j][k][l]){
                            res+=grids[j][k][l];
                        }
                    }
                }                
                res *= numbers[i];
                i = countNumbers;
                break;
            }
        }
    }    
    free(numbers);
    freeLines(lines,size);
    freeLines(numbersString,countNumbers);
    for (int i=0;i<countNumbers;i++){
        for (int j=0;j<sizeGrid;j++){
            free(grids[i][j]);
            free(seen[i][j]);
        }
        free(grids[i]);
        free(seen[i]);
    }
    free(grids);
    free(seen);
    return res;
}

int part2(const char* path){
    int size;
    char **lines = splitFile(path, '\n',&size,true);
    int countNumbers;    
    char **numbersString = splitString(lines[0],',',&countNumbers,true);
    int *numbers = atoiArray(numbersString,countNumbers);
    int countGrid;
    int ***grids = parseGrids(lines, size, &countGrid);
    bool ***seen = initSeen(countGrid);
    int res = 0;
    int lastChecked;
    for (int i=0;i<countNumbers;i++){
        bool checkAll = true;
        for (int j=0;j<countGrid;j++){
            for (int k=0;k<sizeGrid;k++){
                for (int l=0;l<sizeGrid;l++){
                    if (grids[j][k][l] == numbers[i]){
                        seen[j][k][l] = true;
                    }
                }
            }
            if (!checkGrid(seen[j])){                
                checkAll = false;
                lastChecked = j;
            }
        }
        if (checkAll){
            for (int k=0;k<sizeGrid;k++){
                for (int l=0;l<sizeGrid;l++){
                    if (!seen[lastChecked][k][l]){
                        res+=grids[lastChecked][k][l];
                    }
                }
            }            
            res *= numbers[i];            
            break;
        }
    }    
    free(numbers);
    freeLines(lines,size);
    freeLines(numbersString,countNumbers);
    for (int i=0;i<countNumbers;i++){
        for (int j=0;j<sizeGrid;j++){
            free(grids[i][j]);
            free(seen[i][j]);
        }
        free(grids[i]);
        free(seen[i]);
    }
    free(grids);
    free(seen);
    return res;
}
