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
    Lifo *res = createLifo();
    Node *head = lifo->head;
    res->head = head;
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

int getNumberPaths1(Connection *connections, int sizeConnections, Lifo *path){
    char *current = *(char**)getHeadLifo(path);    
    if (strcmp(current, "end") == 0){
        return 1;
    }
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
            if (strcmp(nextCave, "start") == 0){
                Node *head = path->head;
                while (head != NULL){
                    char *step = *(char**)head->elem;
                    if (strcmp(step,"start") == 0){
                        add = false;
                    }
                    head = head->next;
                }
            }
            if (strcmp(nextCave, "end") == 0){
                Node *head = path->head;
                while (head != NULL){
                    char *step = *(char**)head->elem;
                    if (strcmp(step,"end") == 0){
                        add = false;
                    }
                    head = head->next;
                }
            }
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
                Lifo *nextPath = copyLifoString(path);
                addLifo(nextPath, &nextCave);
                res += getNumberPaths1(connections, sizeConnections, nextPath);
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
    for (int i=0;i<size;i++){
        if (strcmp(connections[i].cave1, "start") == 0){
            addLifo(first, &connections[i].cave1);
            break;
        }
        if (strcmp(connections[i].cave2, "start") == 0){
            addLifo(first, &connections[i].cave2);
            break;
        }
    }
    return getNumberPaths1(connections, size, first);
}

int getNumberPaths2(Connection *connections, int sizeConnections, Lifo *path, bool smallCaveTwice){
    char *current = *(char**)getHeadLifo(path);    
    if (strcmp(current, "end") == 0){
        return 1;
    }
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
            bool add1 = true;
            if (strcmp(nextCave, "start") == 0){
                Node *head = path->head;
                while (head != NULL){
                    char *step = *(char**)head->elem;
                    if (strcmp(step,"start") == 0){
                        add1 = false;
                    }
                    head = head->next;
                }
            }
            if (strcmp(nextCave, "end") == 0){
                Node *head = path->head;
                while (head != NULL){
                    char *step = *(char**)head->elem;
                    if (strcmp(step,"end") == 0){
                        add1 = false;
                    }
                    head = head->next;
                }
            }
            bool add2 = true;
            if (nextCave[0] >= 'a' && nextCave[0] <= 'z'){
                Node *head = path->head;                
                while (head != NULL){
                    char *step = *(char**)head->elem;
                    if (strcmp(step,nextCave) == 0){
                        add2 = false;
                    }
                    head = head->next;
                }
            }
            if (add1){
                if (add2){
                    Lifo *nextPath = copyLifoString(path);
                    addLifo(nextPath, &nextCave);
                    res += getNumberPaths2(connections, sizeConnections, nextPath, smallCaveTwice);
                }
                if (!smallCaveTwice && !add2){
                    Lifo *nextPath = copyLifoString(path);
                    addLifo(nextPath, &nextCave);
                    res += getNumberPaths2(connections, sizeConnections, nextPath, true);
                }
            }
        }
    }
    return res;
}

int part2(const char* path){    
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
    for (int i=0;i<size;i++){
        if (strcmp(connections[i].cave1, "start") == 0){
            addLifo(first, &connections[i].cave1);
            break;
        }
        if (strcmp(connections[i].cave2, "start") == 0){
            addLifo(first, &connections[i].cave2);
            break;
        }
    }
    return getNumberPaths2(connections, size, first, false);
}
