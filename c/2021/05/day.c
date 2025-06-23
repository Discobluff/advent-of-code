#include "../../utils/parse.h"
#include "../../utils/point.h"
#include "day.h"
#include <stdio.h>

int max(int a, int b){
    if (a>=b){
        return a;
    }
    return b;
}

int sgn(int a){
    if (a == 0){
        return 0;
    }
    if (a > 0){
        return 1;
    }
    return -1;
}

int result(const char* path, bool part2){
    int size;
    char **lines = splitFile(path, '\n', &size, true);
    int maxX = 0;
    int maxY = 0;
    for (int line=0;line<size;line++){
        int x1,x2,y1,y2;
        sscanf(lines[line],"%d,%d -> %d,%d", &x1, &y1, &x2, &y2);
        maxX = max(maxX,x1);
        maxX = max(maxX,x2);
        maxY = max(maxY,y1);
        maxY = max(maxY,y2);
    }
    int **grid = malloc(sizeof(int*)*(maxY+1));
    for (int y=0;y<=maxY;y++){
        grid[y] = calloc(maxX+1,sizeof(int));
    }
    for (int line=0;line<size;line++){
        int x1,x2,y1,y2;
        sscanf(lines[line],"%d,%d -> %d,%d", &x1, &y1, &x2, &y2);        
        Point p1 = initPoint(x1,y1);
        Point p2 = initPoint(x2,y2);
        Point direction = soustrPoints(p2,p1);
        if (part2 || direction.x == 0 || direction.y == 0){
            if (direction.x != 0){
                direction.x = sgn(direction.x);
            }
            if (direction.y != 0){
                direction.y = sgn(direction.y);
            }            
            while (!equal(p1,p2)){
                grid[p1.y][p1.x] +=1 ;
                p1 = addPoints(p1, direction);
            }
            grid[p2.y][p1.x] +=1;
        }
    }
    
    int res = 0;
    for (int line=0;line<=maxY;line++){
        for (int column=0;column<=maxX;column++){            
            if (grid[line][column] > 1){
                res ++;
            }
        }        
        free(grid[line]);
    }
    free(grid);
    freeLines(lines, size);
    return res;
}

int part1(const char* path){    
    return result(path,false);
}

int part2(const char* path){    
    return result(path,true);
}
