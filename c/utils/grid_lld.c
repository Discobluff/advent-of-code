#include "grid_lld.h"
#include <assert.h>
#include <stdlib.h>

bool isEmpty_lld(Grid_lld *grid){
    return grid->dimension == 0;
}
Grid_lld *initGrid_lld(int dimension, int *lengths, long long int valueDefault){
    assert(dimension > 0);
    Grid_lld *grid = malloc(sizeof(Grid_lld));
    grid->dimension = dimension;
    grid->length = lengths[0];
    if (dimension == 1){
        grid->values = NULL;
        grid->valuesInt = malloc(sizeof(long long int)*grid->length);
        for (int i=0;i<grid->length;i++){
            grid->valuesInt[i] = valueDefault;
        }
        return grid;
    }
    grid->valuesInt = NULL;
    grid->values = malloc(sizeof(Grid_lld*)*grid->length);
    for (int i=0;i<grid->length;i++){
        grid->values[i] = initGrid_lld(dimension-1, lengths+1, valueDefault);
    }
    return grid;
}

long long int getGrid_lld(Grid_lld *grid, int *indexes){
    if (grid->dimension == 1){
        return grid->valuesInt[indexes[0]];
    }
    return getGrid_lld(grid->values[indexes[0]],indexes+1);
}

void setGrid_lld(Grid_lld *grid, int *indexes, long long int value){
    if (grid->dimension == 1){
        grid->valuesInt[indexes[0]] = value;
        return;
    }
    setGrid_lld(grid->values[indexes[0]],indexes+1, value);
}

void freeGrid_lld(Grid_lld *grid){
    if (grid->dimension == 1){
        free(grid->valuesInt);
    }
    else{
        for (int i=0;i<grid->length; i++){
            freeGrid_lld(grid->values[i]);
        }
        free(grid->values);
    }
    free(grid);
}
