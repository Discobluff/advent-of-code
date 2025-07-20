#include "../../utils/parse.h"
#include "../../utils/set.h"
#include "day.h"
#include <stdio.h>
#include <string.h>

char *zero = "abcefg";
char *one = "cf";
char *two = "acdeg";
char *three = "acdfg";
char *four = "bcdf";
char *five = "abdfg";
char *six = "abdefg";
char *seven = "acf";
char *eight = "abcdefg";
char *nine = "abcdfg";
char* numbers[10] = {"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"};

bool checkPart1(char *letters){    
    unsigned long len = strlen(letters);
    return len == strlen(one) || len == strlen(four) || len == strlen(seven) || len == strlen(eight);
}

int part1(const char* path){
    int size;
    char **lines = splitFile(path, '\n', &size, true);
    int res = 0;
    for (int i=0;i<size;i++){
        int sizeLine;
        char **line = splitString(lines[i], '|', &sizeLine, true);
        int nbLetters;
        char **letters = splitString(line[1], ' ', &nbLetters, true);
        for (int j=1;j<nbLetters;j++){
            if (checkPart1(letters[j])){
                res++;
            }
        }
        freeLines(letters, nbLetters);
        freeLines(line, sizeLine);
    }
    freeLines(lines, size);
    return res;
}

bool comp(void *e1, void *e2){
    return *(char*)e1 == *(char*)e2;
}

void printSet(Set *s){
    Node *head = s->head;
    while (head !=NULL){
        printf("%c - ",*(char*)head->elem);
        head = head->next;
    }
    printf("\n");
}

int part2(const char* path){
    int size;
    char **lines = splitFile(path, '\n', &size, true);
    int totalRes = 0;
    for (int i=0;i<size;i++){
        int sizeLine;
        char **line = splitString(lines[i], '|', &sizeLine, true);
        int nbLetters;
        char **letters = splitString(line[0], ' ', &nbLetters, true);
        int nbDigits = strlen(eight);
        Set **candidats = malloc(sizeof(char**)*nbDigits);
        for (int j = 0; j < nbDigits; j++){
            candidats[j] = createSet();
            for (int k = 0; k < nbDigits; k++){
                addSet(candidats[j], &eight[k]);
            }
        }
        Set *setEight = createSet();
        for (int k=0; k<nbDigits;k++){
            addSet(setEight, &eight[k]);
        }
        Set *size6Set = createSet();
        for (int j=0;j<nbLetters;j++){
            int len = strlen(letters[j]);
            if (len == (int)strlen(one)){
                Set *setLetters  = createSet();
                for (int k=0;k<len;k++){
                    addSet(setLetters, &letters[j][k]);
                }
                for (int k=0;k<len;k++){
                    Set *tempSet = intersectSet(setLetters, candidats[one[k]-'a'], comp);
                    freeSet(candidats[one[k]-'a']);
                    candidats[one[k]-'a'] = tempSet;
                }
                freeSet(setLetters);
            }
            if (len == (int)strlen(four)){
                Set *setLetters  = createSet();
                for (int k=0;k<len;k++){
                    addSet(setLetters, &letters[j][k]);
                }
                for (int k=0;k<len;k++){
                    Set *tempSet = intersectSet(setLetters, candidats[four[k]-'a'], comp);
                    freeSet(candidats[four[k]-'a']);
                    candidats[four[k]-'a'] = tempSet;
                }
                freeSet(setLetters);
            }
            if (len == (int)strlen(seven)){
                Set *setLetters  = createSet();
                for (int k=0;k<len;k++){
                    addSet(setLetters, &letters[j][k]);
                }
                for (int k=0;k<len;k++){
                    Set *tempSet = intersectSet(setLetters, candidats[seven[k]-'a'], comp);
                    freeSet(candidats[seven[k]-'a']);
                    candidats[seven[k]-'a'] = tempSet;
                }
                freeSet(setLetters);
            }
            if (len == (int)strlen(eight)){
                Set *setLetters  = createSet();
                for (int k=0;k<len;k++){
                    addSet(setLetters, &letters[j][k]);
                }
                for (int k=0;k<len;k++){
                    Set *tempSet = intersectSet(setLetters, candidats[eight[k]-'a'], comp);
                    freeSet(candidats[eight[k]-'a']);
                    candidats[eight[k]-'a'] = tempSet;
                }
                freeSet(setLetters);
            }
            if (len == (int)strlen(zero)){
                Set *setLetters = createSet();
                for (int k=0;k<len;k++){
                    addSet(setLetters, &letters[j][k]);
                }
                Set *tempSet = priveSet(setEight, setLetters, comp);
                Set *tempTempSet = unionSet(size6Set, tempSet, comp);
                freeSet(size6Set);
                size6Set = tempTempSet;
                freeSet(tempSet);
                freeSet(setLetters);
            }
        }
        Set *tempSet = intersectSet(candidats['d' - 'a'], size6Set, comp);
        freeSet(candidats['d'-'a']);
        candidats['d'-'a'] = tempSet;
        tempSet = intersectSet(candidats['e' - 'a'], size6Set, comp);
        freeSet(candidats['e'-'a']);
        candidats['e'-'a'] = tempSet;
        tempSet = intersectSet(candidats['c' - 'a'], size6Set, comp);
        freeSet(candidats['c'-'a']);
        candidats['c'-'a'] = tempSet;                
        bool go = true;
        while (go){
            go = false;
            for (int j=0;j<nbDigits;j++){
                if (lenSet(candidats[j]) == 1){
                    for (int k=0;k<nbDigits;k++){
                        if (k != j){
                            Set *tempSet = priveSet(candidats[k], candidats[j], comp);
                            freeSet(candidats[k]);
                            candidats[k] = tempSet;
                        }                    
                    }
                }
            }
            for (int j=0;j<nbDigits;j++){
                if (lenSet(candidats[j]) != 1){
                    go = true;
                    break;
                }
            }
        }        
        int nbDecode;
        char **code = splitString(line[1], ' ', &nbDecode, true);
        int res = 0;
        for (int j=0;j<nbDecode;j++){
            Set *lettersRes = createSet();
            for (int k=0;k<(int)strlen(code[j]);k++){
                for (int l=0;l<nbDigits;l++){
                    if (isPresentSet(candidats[l], &code[j][k], comp)){
                        char *letter = malloc(sizeof(char));
                        *letter = l + 'a';
                        addSet(lettersRes, letter);
                    }                    
                }
            }            
            for (int k=0;k<10;k++){
                bool found = true;
                if ((int)strlen(numbers[k]) == lenSet(lettersRes)){
                    
                    for (int l=0;l<(int)strlen(numbers[k]);l++){
                        if (!isPresentSet(lettersRes, &numbers[k][l], comp)){
                            found = false;
                        }
                    }
                    if (found){                        
                        res = 10*res+k;
                        break;
                    }
                }
            }
            freeSetElem(lettersRes);

        }

        for (int j=0;j<nbDigits;j++){
            freeSet(candidats[j]);
        }
        free(candidats);
        freeSet(setEight);
        freeSet(size6Set);
        freeLines(code, nbDecode);
        freeLines(letters, nbLetters);
        freeLines(line, sizeLine);
        totalRes += res;
    }
    freeLines(lines, size);
    return totalRes;
}
