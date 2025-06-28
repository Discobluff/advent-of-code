#include <stdbool.h>

struct _Grid_lld {
    int dimension;
    int length;
    struct _Grid_lld **values;
    long long int *valuesInt;
};

typedef struct _Grid_lld Grid_lld;

bool isEmpty_lld(Grid_lld *grid);
Grid_lld *initGrid_lld(int dimension, int *lengths, long long int valueDefault);
long long int getGrid_lld(Grid_lld *grid, int *indexes);
void setGrid_lld(Grid_lld *grid, int *indexes, long long int value);
void freeGrid_lld(Grid_lld *grid);
