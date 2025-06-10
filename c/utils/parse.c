#include "parse.h"
#include <string.h>

char *parseFile(const char* path){
    FILE* file = fopen(path, "r");
    if (!file) {
        perror("Error opening file");
        return NULL;
    }

    fseek(file, 0, SEEK_END);
    long length = ftell(file);
    fseek(file, 0, SEEK_SET);
    
    char* content = malloc(length + 1);
    if (!content) {
        perror("Error allocating memory");
        fclose(file);
        return NULL;
    }

    fread(content, 1, length, file);
    content[length] = '\0';

    fclose(file);
    return content;
}

char **splitFile(const char* path, char charSplit, int* size, bool ignoreEnd){
    char *file = parseFile(path);
    if (file == NULL){
        return NULL;
    }
    int countLines = 1;
    for (int i=0;file[i]!='\0';i++){
        if (file[i] == charSplit){
            countLines++;
        }
    }
    if (ignoreEnd){
        int countEnd = 0;
        while (strlen(file)-countEnd-1 >= 0 && file[strlen(file)-countEnd-1] == charSplit){
            countEnd++;
            countLines --;
        }
    }
    char **res = (char**)malloc(sizeof(char*)*countLines);
    int indexLine = 0;
    int indexChar = 0;
    while(indexLine < countLines){
        int sizeSequence = 0;
        while (file[indexChar+sizeSequence] != '\0' && file[indexChar+sizeSequence] != charSplit){
            sizeSequence++;
        }
        res[indexLine] = (char*)malloc(sizeof(char)*(sizeSequence+1));
        int indexCurrentLine = 0;
        while (file[indexChar] != '\0' && file[indexChar] != charSplit){            
            res[indexLine][indexCurrentLine] = file[indexChar];
            indexChar++;
            indexCurrentLine++;
        }        
        res[indexLine][indexCurrentLine] = '\0';
        indexChar++;
        indexLine++;
    }
    *size = countLines;    
    return res;

}

int *atoiArray(char **array, int size){
    int *res = (int*)malloc(sizeof(int)*size);
    for (int i=0;i<size;i++){
        res[i] = atoi(array[i]);
    }
    return res;
}

int *splitFileToI(const char* path, char charSplit, int *size, bool ignoreEnd){
    char **lines = splitFile(path, charSplit, size, ignoreEnd);    
    int *res = atoiArray(lines, *size);
    free(lines);
    return res;
}
