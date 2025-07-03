#ifndef PARCOURS_H
#define PARCOURS_H

#include "set.h"

Set *BFS(void *start, Set *(*neighborsFunc)(void *), bool (*comp)(void *, void*));

#endif
