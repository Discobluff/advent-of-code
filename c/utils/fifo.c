#include "fifo.h"
#include <stdlib.h>
#include <assert.h>
#include <stdio.h>
#include "point.h"

Fifo *createFifo(void){
    Fifo *res = malloc(sizeof(Fifo));
    res->head = NULL;
    return res;
}

bool isEmptyFifo(Fifo *fifo){
    return fifo->head == NULL;
}

void addFifo(Fifo *fifo, void *elem){
    Node *newNode = malloc(sizeof(Node));
    newNode->next = NULL;
    newNode->elem = elem;
    if (isEmptyFifo(fifo)){
        fifo->head = newNode;        
        return;
    }
    Node *node = fifo->head;
    while (node->next != NULL){
        node = node->next;
    }
    node->next = newNode;
    return;
}

void *getHeadFifo(Fifo *fifo){
    assert(fifo->head != NULL);
    return fifo->head->elem;
}

void *popFifo(Fifo *fifo){
    void *res = getHeadFifo(fifo);
    Node *next = fifo->head->next;
    free(fifo->head);
    fifo->head = next;
    return res;
}

void freeFifo(Fifo *fifo){
    Node *node = fifo->head;
    while (node != NULL){
        Node *next = node->next;
        free(node);
        node = next;
    }
    free(fifo);
}

void freeFifoElem(Fifo *fifo){
    Node *node = fifo->head;
    while (node != NULL){
        Node *next = node->next;
        free(node->elem);
        free(node);
        node = next;
    }
    free(fifo);
}

void printFifoPoint(Fifo *fifo){
    Node *head = fifo->head;
    while(head !=NULL){
        Point *p = (Point*)(head->elem);
        printf("%d %d,",p->x, p->y);
        head = head->next;
    }
    printf("\n");
}

int lenFifo(Fifo *fifo){
    Node *head = fifo->head;
    int res = 0;
    while (head !=NULL){
        res++;
        head = head->next;   
    }
    return res;
}
