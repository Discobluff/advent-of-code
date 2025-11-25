
#ifndef FIFO_H
#define FIFO_H

#include <stdbool.h>

struct _Node {
    struct _Node *next;
    void *elem;
};

typedef struct _Node Node;

struct _Fifo {
    Node *head;
};

typedef struct _Fifo Fifo;

Fifo *createFifo(void);
bool isEmptyFifo(Fifo *fifo);
void addFifo(Fifo *fifo, void *elem);
void *getHeadFifo(Fifo *fifo);
void *popFifo(Fifo *fifo);
void freeFifo(Fifo *fifo);
void freeFifoElem(Fifo *fifo);
void printFifoPoint(Fifo *fifo);
int lenFifo(Fifo *fifo);

#endif
