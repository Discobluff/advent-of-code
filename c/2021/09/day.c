#include "../../utils/parse.h"
#include "day.h"
#include <stdio.h>
#include "../../utils/point.h"
#include <string.h>

int part1(const char* path){
    int size;
    int res = 0;
    char ** lines = splitFile(path, '\n', &size, true);
    for (int line = 0;line<size;line++){
        int sizeColumn = strlen(lines[line]);
        for (int column = 0; column < sizeColumn; column++){
            bool lowPoint = true;
            Point p = initPoint(column, line);            
            if (line > 0 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, NORTH_POINT))){
                lowPoint = false;
            }
            if (column > 0 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, WEST_POINT))){
                lowPoint = false;
            }
            if (line < size-1 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, SOUTH_POINT))){
                lowPoint = false;
            }            
            if (column < sizeColumn-1 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, EAST_POINT))){
                lowPoint = false;
            }
            if (lowPoint){                
                res += 1 + EVAL_POINT(lines, p) - '0';
            }
        }
    }
    freeLines(lines, size);
    return res;
}

int part2(const char* path){    
    return 0;
}
