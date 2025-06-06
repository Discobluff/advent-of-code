CC = clang
CFLAGS = -Wall -Wextra -Wpedantic -std=c99 #-fsanitize=address -g

utils/%.o: utils/%.c
	@$(CC) $(CFLAGS) -c $^ -o $@

2021/%/main.o: 2021/%/main.c 
	@$(CC) $(CFLAGS) -o $@ -c $^

2021/%/main: utils/parse.o 2021/%/main.o
	@$(CC) $(CFLAGS) -o $@ $^

.SILENT:
run2021-%: 2021/%/main
	@./$^

clean:
	@rm -rf main *.o