
#ifndef LIFO_H
#define LIFO_H

#include <stdbool.h>

struct _Node {
    struct _Node *next;
    void *elem;
};

typedef struct _Node Node;

struct _Lifo {
    Node *head;
};

typedef struct _Lifo Lifo;

Lifo *createLifo(void);
bool isEmptyLifo(Lifo *lifo);
void addLifo(Lifo *lifo, void *elem);
void *getHeadLifo(Lifo *lifo);
void *popLifo(Lifo *lifo);
void *popLifoFree(Lifo *lifo);
void freeLifo(Lifo *lifo);
void freeLifoElem(Lifo *lifo);
int lenLifo(Lifo *lifo);

#endif
