#ifndef PARSE_H
#define PARSE_H

#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

char *parseFile(const char* path);
int *atoiArray(char **array, int size);
char **splitFile(const char* path, char charSplit, int* size, bool ignoreEnd);
int *splitFileToI(const char* path, char charSplit, int *size, bool ignoreEnd);
char **splitString(char *file, char charSplit, int *size, bool ignoreEnd);
void freeLines(char **lines, int size);

#endif
