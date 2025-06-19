#include "../../utils/parse.h"
#include <stdio.h>
#include <string.h>

int mostCommon(char **numbers, int size, int index){
    int countOne = 0;
    for (int i=0;i<size;i++){
        if (numbers[i][index] == '1'){
            countOne ++;
        }
    }    
    if (countOne > size/2){
        return 1;
    }
    return 0;
}

int part1(const char *path){
    int size;
    char **numbers = splitFile(path,'\n',&size, true);
    int sizeNumber = strlen(numbers[0]);
    int n1 = 0;
    int n2 = 0;
    for (int i=0;i<sizeNumber;i++){
        int bit = mostCommon(numbers,size,i);        
        n1 = 2*n1+bit;
        n2 = 2*n2+1-bit;
    }    
    freeLines(numbers, size);
    return n1*n2;
}

int mostCommonFilter(char **numbers, int size, int index, bool *filter){
    int countOne = 0;
    int sizeFilter = 0;
    for (int i=0;i<size;i++){
        if (filter[i]){
            sizeFilter++;        
            if (numbers[i][index] == '1'){
                countOne ++;
            }
        }
    }    
    if ((float)countOne >= (float)sizeFilter/2){
        return 1;
    }
    return 0;
}

int convertBinToDec(char *n){
    int result = 0;
    for (int i = 0; n[i] != '\0'; i++) {
        result = result * 2 + (n[i] - '0');
    }
    return result;
}

int part2(const char *path){    
    int size;
    char **numbers = splitFile(path,'\n',&size, true);    
    bool *bits1 = malloc(size*sizeof(bool));
    bool *bits2 = malloc(size*sizeof(bool));
    for (int i=0;i<size;i++){
        bits1[i] = true;
        bits2[i] = true;
    }

    int n1 = 0;
    int count1 = size;
    int index1 = 0;
    while (count1 != 1){
        int bit = mostCommonFilter(numbers,size, index1,bits1);        
        for (int i=0;i<size;i++){
            if (bits1[i] && numbers[i][index1] != bit + '0'){
                bits1[i] = false;
                count1--;
            }
        }
        index1 ++;
    }
    for (int i=0;i<size;i++){
        if (bits1[i]){            
            n1 = convertBinToDec(numbers[i]);
            break;
        }
    }
    int n2 = 0;
    int count2 = size;
    int index2 = 0;
    while (count2 != 1){
        int bit = 1-mostCommonFilter(numbers,size, index2,bits2);        
        for (int i=0;i<size;i++){
            if (bits2[i] && numbers[i][index2] != bit + '0'){
                bits2[i] = false;
                count2--;
            }
        }
        index2 ++;
    }
    for (int i=0;i<size;i++){
        if (bits2[i]){            
            n2 = convertBinToDec(numbers[i]);
            break;
        }
    }
    freeLines(numbers,size);
    free(bits1);
    free(bits2);    
    return n1*n2;
}
