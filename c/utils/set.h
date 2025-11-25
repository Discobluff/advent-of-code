#ifndef SET_H
#define SET_H

#include "fifo.h"
#include <stdbool.h>

typedef Fifo Set;

Set *createSet(void);
bool isEmptySet(Set *set);
bool isPresentSet(Set *set, void *elem, bool (*comp)(void *, void*));
void addSet(Set *set, void *elem);
void *minSet(Set *set, bool (*comp)(void *, void*));
void removeSet(Set *set, void * elem, bool (*comp)(void *, void*));
void removeSetFree(Set *set, void * elem, bool (*comp)(void *, void*));
int lenSet(Set *set);
int minSetInt(Set *set);
int sumSet(Set *set);
int prodSet(Set *set);
void freeSetElem(Set *set);
void freeSet(Set *set);
Set *intersectSet(Set *s1, Set *s2, bool (*comp)(void *, void*));
Set *unionSet(Set *s1, Set *s2, bool (*comp)(void *, void*));
Set *priveSet(Set *s1, Set *s2, bool (*comp)(void *, void*));

#endif
