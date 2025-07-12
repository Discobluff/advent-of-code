#include "../../utils/parse.h"
#include "../../utils/lifo.h"
#include "../../utils/sort.h"
#include <string.h>
#include "day.h"
#include <stdio.h>

int part1(const char* path){
    int size;
    char **lines = splitFile(path, '\n', &size, true);
    int score = 0;
    for (int i = 0;i<size;i++){
        Lifo *lifo = createLifo();
        int len = strlen(lines[i]);
        for (int j=0;j<len;j++){
            char ch = lines[i][j]; 
            char head;
            if (isEmptyLifo(lifo)){
                head = 0;
            }
            else{
                head = *(char*)getHeadLifo(lifo);
            }            
            switch (ch){
                case '[':
                case '(':
                case '{':
                case '<':{
                    char *new = (char*)malloc(sizeof(char));
                    *new = ch;
                    addLifo(lifo, new);
                    break;
                }
                case ')':
                    if (head != '('){
                        score += 3;                        
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;
                case ']':
                    if (*(char*)getHeadLifo(lifo) != '['){
                        score += 57;                        
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;
                case '}':
                    if (*(char*)getHeadLifo(lifo) != '{'){
                        score += 1197;                        
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;                
                case '>':
                    if (*(char*)getHeadLifo(lifo) != '<'){
                        score += 25137;                        
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;
            }
        }
        freeLifo(lifo);
    }
    freeLines(lines, size);
    return score;
}

long part2(const char* path){
    int size;
    char **lines = splitFile(path, '\n', &size, true);
    long *scores = malloc(sizeof(long)*size);
    int lenScores = 0;
    for (int i = 0;i<size;i++){
        Lifo *lifo = createLifo();
        int len = strlen(lines[i]);
        bool incorrect = false;
        for (int j=0;j<len;j++){
            char ch = lines[i][j]; 
            char head;
            if (isEmptyLifo(lifo)){
                head = 0;
            }
            else{
                head = *(char*)getHeadLifo(lifo);
            }            
            switch (ch){
                case '[':
                case '(':
                case '{':
                case '<':{
                    char *new = (char*)malloc(sizeof(char));
                    *new = ch;
                    addLifo(lifo, new);
                    break;
                }
                case ')':
                    if (head != '('){
                        incorrect = true;
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;
                case ']':
                    if (*(char*)getHeadLifo(lifo) != '['){
                        incorrect = true;
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;
                case '}':
                    if (*(char*)getHeadLifo(lifo) != '{'){
                        incorrect = true;
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;                
                case '>':
                    if (*(char*)getHeadLifo(lifo) != '<'){
                        incorrect = true;
                        j = len;
                    }
                    else{
                        popLifo(lifo);
                    }
                    break;
            }
        }
        if (!incorrect){
            long score = 0;
            Node *node = lifo->head;
            while(node != NULL){
                score *= 5;
                char ch = *(char*)node->elem;
                switch (ch){
                    case '(':
                        score += 1;
                        break;
                    case '[':
                        score += 2;
                        break;
                    case '{':
                        score += 3;
                        break;
                    case '<':
                        score += 4;
                        break;
                }
                node = node->next;
            }
            scores[lenScores] = score;
            lenScores++;
        }
        freeLifo(lifo);
    }
    sortArrayLong(scores, lenScores);
    freeLines(lines, size);
    return scores[lenScores/2];
}
