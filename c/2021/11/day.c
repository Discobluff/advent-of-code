#include "../../utils/parse.h"
#include "day.h"
#include <stdio.h>
#include "../../utils/point.h"
#include <string.h>

bool validPosition(Point pos, int sizeLine, int sizeColumn){
    return pos.x >= 0 && pos.y >= 0 && pos.y < sizeLine && pos.x < sizeColumn;
}

void flash(int **grid, int sizeLine, int sizeColumn, Point pos){
    for (int i=0;i<LEN_CARDINAL_POINTS;i++){
        Point newPoint = addPoints(pos, CARDINAL_POINTS[i]);
        if (validPosition(newPoint, sizeLine, sizeColumn)){
            if (grid[newPoint.y][newPoint.x] != 0){
                grid[newPoint.y][newPoint.x] += 1;
            }
        }
    }
    grid[pos.y][pos.x] = 0;
}

int evolveStep(int **grid, int sizeLine, int sizeColumn){
    int res = 0;
    for (int line=0;line<sizeLine;line++){        
        for (int column=0;column<sizeColumn;column++){
            grid[line][column] += 1;            
        }
    }
    bool go = true;
    while (go){
        go = false;
        for (int line=0;line<sizeLine;line++){        
            for (int column=0;column<sizeColumn;column++){
                if (grid[line][column] > 9){
                    Point pos = initPoint(column, line);
                    flash(grid, sizeLine, sizeColumn, pos);
                    res++;
                    go = true;
                }
            }
        }
    }
    return res;
}

int part1(const char* path){
    int size;
    int res = 0;
    char **lines = splitFile(path, '\n', &size, true);
    int len = strlen(lines[0]);
    int **grid = malloc(sizeof(int*)*size);
    for (int i=0;i<size;i++){
        grid[i] = malloc(sizeof(int)*len);
        for (int j=0;j<len;j++){
            grid[i][j] = lines[i][j] - '0';
        }
    }
    freeLines(lines, size);
    int countStep = 100;
    for (int i=0;i<countStep;i++){
        res += evolveStep(grid, size, len);
    }
    for (int i=0;i<size;i++){
        free(grid[i]);
    }
    free(grid);
    return res;
}

int part2(const char* path){    
    int size;    
    char **lines = splitFile(path, '\n', &size, true);
    int len = strlen(lines[0]);
    int **grid = malloc(sizeof(int*)*size);
    for (int i=0;i<size;i++){
        grid[i] = malloc(sizeof(int)*len);
        for (int j=0;j<len;j++){
            grid[i][j] = lines[i][j] - '0';
        }
    }
    freeLines(lines, size);
    int countStep = 0;
    int res = 0;
    while (res ==0){
        if (evolveStep(grid, size, len) == len * size){
            res = countStep+1;            
        }
        countStep++;
    }
    for (int i=0;i<size;i++){
        free(grid[i]);
    }
    free(grid);
    return res;
}
