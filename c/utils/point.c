#include "point.h"

Point addPoints(Point p1, Point p2){
    Point p3;
    p3.x = p1.x + p2.x;
    p3.y = p1.y + p2.y;
    return p3;
}
Point initPoint(int x, int y){
    Point p = {.x = x, .y = y};
    return p;
}

Point opposePoint(Point p){
    Point res;
    res.x = -p.x;
    res.y = -p.y;
    return res;
}

Point soustrPoints(Point p1, Point p2){
    return addPoints(p1, opposePoint(p2));
}

bool equal(Point p1, Point p2){
    return p1.x == p2.x && p1.y == p2.y;
}

Point divScalPoint(Point p, int s){
    Point res;
    res.x = p.x/s;
    res.y = p.y/s;
    return res;
}
