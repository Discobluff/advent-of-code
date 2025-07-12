#include "set.h"
#include <stdlib.h>
#include <assert.h>
#include <stdio.h>

Set *createSet(void){
    return createFifo();
}

bool isEmptySet(Set *set){
    return isEmptyFifo(set);
}

bool isPresentSet(Set *set, void *elem, bool (*comp)(void *, void*)){
    Node *node = set->head;
    while (node != NULL){
        if (comp(node->elem, elem)){
            return true;
        }
        node = node->next;
    }
    return false;
}

void addSet(Set *set, void *elem){
    addFifo(set, elem);
}

void *minSet(Set *set, bool (*comp)(void *, void*)){ // arg1 <= arg2
    assert(!isEmptySet(set));
    void *res = set->head->elem;
    Node *head = set->head;
    while (head != NULL){
        if (comp(head->elem, res)){
            res = head->elem;
        }
        head = head->next;
    }
    return res;
}

int minSetInt(Set *set){
    assert(!isEmptySet(set));
    Node *head = set->head;
    int res = *(int*)head->elem;
    while(head !=NULL){
        int elem = *(int*)head->elem;
        if (elem <= res){
            res = elem;
        }
        head = head->next;
    }
    return res;
}

void removeSet(Set *set, void * elem, bool (*comp)(void *, void*)){
    Node *head = set->head;
    if (comp(head->elem, elem)){
        Node *next = head->next;
        free(head);
        set->head = next;
        return;
    }
    while (head->next != NULL){
        if (comp(head->next->elem, elem)){
            Node *next = head->next->next;
            free(head->next);
            head->next = next;
            return;
        }
        head = head->next;
    }
}

void removeSetFree(Set *set, void * elem, bool (*comp)(void *, void*)){
    Node *head = set->head;
    if (comp(head->elem, elem)){
        Node *next = head->next;
        free(head->elem);
        free(head);
        set->head = next;
        return;
    }
    while (head->next != NULL){
        if (comp(head->next->elem, elem)){
            Node *next = head->next->next;
            free(head->next->elem);
            free(head->next);
            head->next = next;
            return;
        }
        head = head->next;
    }
}

int lenSet(Set *set){
    Node *head = set->head;
    int len = 0;
    while (head != NULL){
        len++;
        head = head->next;
    }
    return len;
}

int sumSet(Set *set){
    int res = 0;
    Node *head = set->head;
    while(head !=NULL){
        res += *(int*)head->elem;
        head = head->next;
    }
    return res;
}

int prodSet(Set *set){
    int res = 1;
    Node *head = set->head;
    while(head !=NULL){        
        res *= *(int*)head->elem;
        head = head->next;
    }
    return res;
}

void freeSetElem(Set *set){
    freeFifoElem(set);
}

void freeSet(Set *set){
    freeFifo(set);
}
