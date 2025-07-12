#ifndef POINT_H
#define POINT_H

#include <stdbool.h>

struct _Point{
    int x;
    int y;
};

typedef struct _Point Point;

#define NORTH_POINT ((Point){.x = 0, .y = -1})
#define SOUTH_POINT ((Point){.x = 0, .y = 1})
#define EAST_POINT ((Point){.x = 1, .y = 0})
#define WEST_POINT ((Point){.x = -1, .y = 0})
#define CENTER_POINT ((Point){.x = 0, .y = 0})
#define NORTH_WEST_POINT ((Point){.x = -1, .y = -1})
#define SOUTH_WEST_POINT ((Point){.x = -1, .y = 1})
#define NORTH_EAST_POINT ((Point){.x = 1, .y = -1})
#define SOUTH_EAST_POINT ((Point){.x = 1, .y = 1})
#define CENTER_POINT ((Point){.x = 0, .y = 0})

#define EVAL_POINT(t, p) (t[p.y][p.x])

#define LEN_CARDINAL_POINTS 8

static const Point CARDINAL_POINTS[] = {
    NORTH_POINT,
    SOUTH_POINT,
    EAST_POINT,
    WEST_POINT,
    NORTH_WEST_POINT,
    SOUTH_WEST_POINT,
    NORTH_EAST_POINT,
    SOUTH_EAST_POINT
};

Point addPoints(Point p1, Point p2);
Point initPoint(int x, int y);
Point opposePoint(Point p);
Point soustrPoints(Point p1, Point p2);
bool equal(Point p1, Point p2);
Point divScalPoint(Point p, int s);
Point *mallocPoint(Point p);

#endif
