#include <stdbool.h>

struct _Grid {
    int dimension;
    int length;
    struct _Grid **values;
    long long int *valuesInt;
};

typedef struct _Grid Grid;

bool isEmpty(Grid *grid);
Grid *initGrid(int dimension, int *lengths, int valueDefault);
int getGrid(Grid *grid, int *indexes);
void setGrid(Grid *grid, int *indexes, int value);
void freeGrid(Grid *grid);
