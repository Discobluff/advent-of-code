#include "lifo.h"
#include <stdlib.h>
#include <assert.h>
#include <stdio.h>
#include "point.h"

Lifo *createLifo(void){
    Lifo *res = malloc(sizeof(Lifo));
    res->head = NULL;
    return res;
}

bool isEmptyLifo(Lifo *lifo){
    return lifo->head == NULL;
}

void addLifo(Lifo *lifo, void *elem){
    Node *newNode = malloc(sizeof(Node));
    newNode->next = lifo->head;
    newNode->elem = elem;
    lifo->head = newNode;
}

void *getHeadLifo(Lifo *lifo){
    assert(lifo->head != NULL);
    return lifo->head->elem;
}

void *popLifo(Lifo *lifo){
    void *res = getHeadLifo(lifo);
    Node *next = lifo->head->next;
    free(lifo->head);
    lifo->head = next;
    return res;
}

void freeLifo(Lifo *lifo){
    Node *node = lifo->head;
    while (node != NULL){
        Node *next = node->next;
        free(node);
        node = next;
    }
    free(lifo);
}

int lenLifo(Lifo *lifo){
    Node *head = lifo->head;
    int res = 0;
    while (head !=NULL){
        res++;
        head = head->next;   
    }
    return res;
}
