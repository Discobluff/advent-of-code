#include <stdio.h>
#include <string.h>
#include "day.h"
#include "../../utils/parse.h"
#include "../../utils/point.h"
#include "../../utils/set.h"
#include "../../utils/parcours.h"

struct _NodeBFS{
    Point p;
    int size;
    char **lines;
};

typedef struct _NodeBFS NodeBFS;

int part1(const char *path){
    int size;
    int res = 0;
    char **lines = splitFile(path, '\n', &size, true);
    for (int line = 0; line < size; line++)
    {
        int sizeColumn = strlen(lines[line]);
        for (int column = 0; column < sizeColumn; column++)
        {
            bool lowPoint = true;
            Point p = initPoint(column, line);
            if (line > 0 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, NORTH_POINT)))
            {
                lowPoint = false;
            }
            if (column > 0 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, WEST_POINT)))
            {
                lowPoint = false;
            }
            if (line < size - 1 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, SOUTH_POINT)))
            {
                lowPoint = false;
            }
            if (column < sizeColumn - 1 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, EAST_POINT)))
            {
                lowPoint = false;
            }
            if (lowPoint)
            {
                res += 1 + EVAL_POINT(lines, p) - '0';
            }
        }
    }
    freeLines(lines, size);
    return res;
}


Set *getNeighbors(void *point) {
    Set *neighbors = createSet();
    NodeBFS *p = (NodeBFS *)point;
    // printf("%d %d\n",p->p.x, p->p.y);
    int size = p->size;
    char **lines = p->lines;
    Point directions[] = {NORTH_POINT, SOUTH_POINT, EAST_POINT, WEST_POINT};
    for (int i = 0; i < 4; i++) {
        Point neighbor = addPoints(p->p, directions[i]);
        // printf("%d %d\n",neighbor.x, neighbor.y);
        if (neighbor.y >= 0 && neighbor.y < size && neighbor.x >= 0 && neighbor.x < (int)strlen(lines[neighbor.y])) {
            if (EVAL_POINT(lines, neighbor) != '9') {
                NodeBFS *neighborNode = malloc(sizeof(NodeBFS));
                neighborNode->lines = lines;
                neighborNode->size = size;
                neighborNode->p = neighbor;
                addSet(neighbors, neighborNode);
            }
        }
    }

    return neighbors;
}

bool compNodeBFS(void *e1, void *e2){
    NodeBFS *e3 = (NodeBFS*)e1;
    NodeBFS *e4 = (NodeBFS*)e2;
    return e3->p.x == e4->p.x && e3->p.y == e4->p.y;
}

bool compRemoveSet(void *e1, void *e2){
    return *(int*)e1 == *(int*)e2;
}

int part2(const char *path){
    int size;
    Set *res = createSet();
    char **lines = splitFile(path, '\n', &size, true);
    for (int line = 0; line < size; line++)
    {
        int sizeColumn = strlen(lines[line]);
        for (int column = 0; column < sizeColumn; column++)
        {
            bool lowPoint = true;
            Point p = initPoint(column, line);
            if (line > 0 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, NORTH_POINT)))
            {
                lowPoint = false;
            }
            if (column > 0 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, WEST_POINT)))
            {
                lowPoint = false;
            }
            if (line < size - 1 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, SOUTH_POINT)))
            {
                lowPoint = false;
            }
            if (column < sizeColumn - 1 && EVAL_POINT(lines, p) >= EVAL_POINT(lines, addPoints(p, EAST_POINT)))
            {
                lowPoint = false;
            }
            if (lowPoint){
                NodeBFS pointNode;
                pointNode.lines = lines;
                pointNode.size = size;
                pointNode.p = p;
                Set *basin = BFS(&pointNode, getNeighbors, compNodeBFS);
                int *len = malloc(sizeof(int));
                *len = lenSet(basin);                
                if (lenSet(res) < 3){
                    addSet(res, len);
                }
                else{
                    int min = minSetInt(res);                    
                    if (min < *len){
                        removeSet(res, &min,compRemoveSet);
                        addSet(res, len);
                    }
                }                
            }
        }
    }
    freeLines(lines, size);
    return prodSet(res);
}
