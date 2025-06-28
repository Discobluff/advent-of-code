#include "grid.h"
#include <assert.h>
#include <stdlib.h>

bool isEmpty(Grid *grid){
    return grid->dimension == 0;
}
Grid *initGrid(int dimension, int *lengths, int valueDefault){
    assert(dimension > 0);
    Grid *grid = malloc(sizeof(Grid));
    grid->dimension = dimension;
    grid->length = lengths[0];
    if (dimension == 1){
        grid->values = NULL;
        grid->valuesInt = malloc(sizeof(int)*grid->length);
        for (int i=0;i<grid->length;i++){
            grid->valuesInt[i] = valueDefault;
        }
        return grid;
    }
    grid->valuesInt = NULL;
    grid->values = malloc(sizeof(Grid*)*grid->length);
    for (int i=0;i<grid->length;i++){
        grid->values[i] = initGrid(dimension-1, lengths+1, valueDefault);
    }
    return grid;
}

int getGrid(Grid *grid, int *indexes){
    if (grid->dimension == 1){
        return grid->valuesInt[indexes[0]];
    }
    return getGrid(grid->values[indexes[0]],indexes+1);
}

void setGrid(Grid *grid, int *indexes, int value){
    if (grid->dimension == 1){
        grid->valuesInt[indexes[0]] = value;
        return;
    }
    setGrid(grid->values[indexes[0]],indexes+1, value);
}

void freeGrid(Grid *grid){
    if (grid->dimension == 1){
        free(grid->valuesInt);
    }
    else{
        for (int i=0;i<grid->length; i++){
            freeGrid(grid->values[i]);
        }
        free(grid->values);
    }
    free(grid);
}
