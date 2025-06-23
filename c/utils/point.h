#include <stdbool.h>

struct _Point{
    int x;
    int y;
};

typedef struct _Point Point;

#define NORTH_POINT ({.x = 0, .y = -1})
#define SOUTH_POINT ({.x = 0; .y = 1})
#define EAST_POINT ({.x = -1; .y = 0})
#define WEST_POINT ({.x = 1; .y = 0})
#define CENTER_POINT ({.x = 0; .y = 0})

Point addPoints(Point p1, Point p2);
Point initPoint(int x, int y);
Point opposePoint(Point p);
Point soustrPoints(Point p1, Point p2);
bool equal(Point p1, Point p2);
Point divScalPoint(Point p, int s);
