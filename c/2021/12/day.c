#include "../../utils/parse.h"
#include "../../utils/lifo.h"
#include "../../utils/fifo.h"
#include "day.h"
#include <stdio.h>
#include <string.h>

struct _Connection{
    char *cave1;
    char *cave2;
};

typedef struct _Connection Connection;

Lifo *copyLifoString(Lifo *lifo){
    Fifo *fifo = createFifo();
    Node *head = lifo->head;
    while (head != NULL){
        char *oldElem = *(char**)head->elem;
        char *newElem = malloc(sizeof(char)*(strlen(oldElem)+1));
        strcpy(newElem, oldElem);
        addFifo(fifo, &newElem);
        head = head->next;
    }
    Lifo *res = createLifo();
    res->head = fifo->head;
    return res;
}

void printLifoString(Lifo *lifo){
    Node *h = lifo->head;
    while(h !=NULL){
        char *string = *(char**)h->elem;
        printf("%s - ",string);
        h = h->next;
    }
    printf("\n");
}

int getNumberPaths(Connection *connections, int sizeConnections, Lifo *path){
    char *current = *(char**)getHeadLifo(path);    
    if (strcmp(current, "end") == 0){
        return 1;
    }
    if (lenLifo(path)>3){
        return 1;
    }
    // printLifoString(path);
    printf("%d\n",lenLifo(path));
    printf("current : %s\n",current);
    int res = 0;
    for (int i=0;i<sizeConnections;i++){
        bool b1 = (strcmp(current, connections[i].cave1) == 0);
        bool b2 = (strcmp(current, connections[i].cave2) == 0);
        char *nextCave;
        if (b1){
            nextCave = connections[i].cave2;
        }
        if (b2){
            nextCave = connections[i].cave1;
        }
        if (b1 || b2){
            bool add = true;
            if (nextCave[0] >= 'a' && nextCave[0] <= 'z'){
                Node *head = path->head;
                while (head != NULL){
                    char *step = *(char**)head->elem;
                    if (strcmp(step,nextCave) == 0){
                        add = false;
                    }
                    head = head->next;
                }
            }
            if (add){
                printf("next : %s\n",nextCave);
                Lifo *nextPath = copyLifoString(path);
                char *newCave = malloc(sizeof(char)*(strlen(nextCave)+1));
                strcpy(newCave, nextCave);
                // printLifoString(nextPath);
                // printf("%s\n",newCave);
                addLifo(nextPath, &newCave);
                // nextPath->head->next = (Node*)malloc(sizeof(Node));
                // nextPath->head->next->elem = newCave;
                // printf("marge\n");
                // printLifoString(nextPath);
                // printLifoString(nextPath);
                res += getNumberPaths(connections, sizeConnections, nextPath);
            }
        }
    }
    return res;
}

int part1(const char* path){
    int size;
    char **lines = splitFile(path, '\n', &size, true);
    Connection *connections = malloc(sizeof(Connection)*size);
    for (int i=0;i<size;i++){
        char *c1;
        char *c2;
        sscanf(lines[i], "%m[^-]-%m[^-]", &c1, &c2);
        connections[i].cave1 = c1;
        connections[i].cave2 = c2;
    }
    freeLines(lines, size);
    Lifo *first = createLifo();
    char *start = malloc(sizeof(char)*6);
    strcpy(start, "start");
    addLifo(first, &start);
    return getNumberPaths(connections, size, first);
}

int part2(const char* path){    
    return 0;
}
