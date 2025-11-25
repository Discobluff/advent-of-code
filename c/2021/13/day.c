#include "../../utils/parse.h"
#include "day.h"
#include <string.h>
#include <stdio.h>

bool **foldY(bool **grid, int sizeX, int sizeY){
    bool **res = malloc(sizeof(bool*)*(sizeY-1)/2);
    for (int i=0;i<(sizeY-1)/2;i++){
        res[i] = calloc(sizeX, sizeof(bool));
        for (int j=0;j<sizeX;j++){
            res[i][j] = res[i][j] || grid[i][j];
        }
    }
    for (int i=(sizeY+1)/2;i<sizeY;i++){
        for (int j=0;j<sizeX;j++){
            res[i-(sizeY+1)/2][j] = res[i-(sizeY+1)/2][j] || grid[sizeY-1-i+(sizeY+1)/2][j];
        }
    }
    return res;
}

bool **foldX(bool **grid, int sizeX, int sizeY){
    bool **res = malloc(sizeof(bool*)*sizeY);
    for (int i=0;i<sizeY;i++){
        res[i] = calloc((sizeX-1)/2, sizeof(bool));
        for (int j=0;j<(sizeX-1)/2;j++){
            res[i][j] = res[i][j] || grid[i][j];
        }
    }
    for (int i=0;i<sizeY;i++){
        for (int j=(sizeX+1)/2;j<sizeX;j++){
            res[i][j-(sizeX+1)/2] = res[i][j-(sizeX+1)/2] || grid[i][sizeX-1-j+(sizeX+1)/2];
        }
    }
    return res;
}

void printGrid(bool **grid, int sizeX, int sizeY){
    for (int i=0;i<sizeY;i++){
        for (int j=0;j<sizeX;j++){
            if (grid[i][j]){
                printf("#");
            }
            else{
                printf(".");
            }
        }
        printf("\n");
    }
}

void freeGrid(bool **grid, int sizeY){
    for (int i=0;i<sizeY;i++){
        free(grid[i]);
    }
    free(grid);
}

int part1(const char* path){
    int size;
    char **input = splitFile(path, '\n', &size, true);
    int maxX, maxY;
    int countCoords = 0;
    for (int i=0; strcmp(input[i],"");i++){
        int x,y;
        sscanf(input[i], "%d,%d", &x, &y);
        if (x > maxX){
            maxX = x;
        }
        if (y > maxY){
            maxY = y;
        }
        countCoords++;
    }
    int sizeX = maxX + 1;
    int sizeY = maxY + 1;
    bool **grid = malloc(sizeof(bool*)*sizeY);
    for (int i=0;i<maxY+1;i++){
        grid[i] = calloc(sizeX,sizeof(bool));
    }
    for (int i=0; strcmp(input[i],"");i++){
        int x,y;
        sscanf(input[i], "%d,%d", &x, &y);
        grid[y][x] = true;
    }
    int coord;
    char c;
    sscanf(input[countCoords+1], "fold along %c=%d", &c, &coord);
    bool **temp;
    if (c == 'x'){
        temp = foldX(grid, sizeX, sizeY);
        freeGrid(grid, sizeY);
        sizeX = (sizeX-1)/2;
    }
    else{
        temp = foldY(grid, sizeX, sizeY);
        freeGrid(grid, sizeY);
        sizeY = (sizeY-1)/2;
    }
    grid = temp;
    int res = 0;
    for (int i=0; i<sizeY;i++){
        for (int j=0;j<sizeX;j++){
            if (grid[i][j]){
                res++;
            }
        }
    }
    freeGrid(grid, sizeY);
    freeLines(input, size);
    return res;
}

int part2(const char* path){    
    int size;
    char **input = splitFile(path, '\n', &size, true);
    int maxX, maxY;
    int countCoords = 0;
    for (int i=0; strcmp(input[i],"");i++){
        int x,y;
        sscanf(input[i], "%d,%d", &x, &y);
        if (x > maxX){
            maxX = x;
        }
        if (y > maxY){
            maxY = y;
        }
        countCoords++;
    }
    int sizeX = maxX + 1;
    int sizeY = maxY + 1;
    bool **grid = malloc(sizeof(bool*)*sizeY);
    for (int i=0;i<maxY+1;i++){
        grid[i] = calloc(sizeX,sizeof(bool));
    }
    for (int i=0; strcmp(input[i],"");i++){
        int x,y;
        sscanf(input[i], "%d,%d", &x, &y);
        grid[y][x] = true;
    }
    for (int i=countCoords+1; i<size;i++){
        int coord;
        char c;
        sscanf(input[i], "fold along %c=%d", &c, &coord);
        bool **temp;
        if (c == 'x'){
            temp = foldX(grid, sizeX, sizeY);
            freeGrid(grid, sizeY);
            sizeX = (sizeX-1)/2;
        }
        else{
            temp = foldY(grid, sizeX, sizeY);
            freeGrid(grid, sizeY);
            sizeY = (sizeY-1)/2;
        }
        grid = temp;
    }
    printGrid(grid, sizeX, sizeY);
    freeGrid(grid, sizeY);
    freeLines(input, size);
    return 0;
}
