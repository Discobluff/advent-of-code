#include "parcours.h"
#include "set.h"
#include "fifo.h"
#include <stdlib.h>

Set *BFS(void *start, Set *(*neighborsFunc)(void *), bool (*comp)(void *, void*)){
    Set *seen = createSet();
    Fifo *nexts = createFifo();
    addFifo(nexts, start);    
    while (!isEmptyFifo(nexts)){        
        void *current = getHeadFifo(nexts);
        if (!isPresentSet(seen, current, comp)){
            Set *neighbors = (neighborsFunc)(current);
            Node *head = neighbors->head;
            while (head != NULL){
                addFifo(nexts, head->elem);
                head = head->next;
            }
            freeSet(neighbors);
            addSet(seen, current);
        }
        else{
            free(current);
        }
        Node *temp = nexts->head->next;
        free(nexts->head);
        nexts->head = temp;
    }    
    freeFifo(nexts);
    return seen;
}
