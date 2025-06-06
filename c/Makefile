CC = clang
CFLAGS = -Wall -Wextra -Wpedantic -std=c99 #-fsanitize=address -g

%.o: utils/%.c
	@$(CC) $(CFLAGS) -c $^ -o $@

2021/%/day%.o: 2021/%/day%.c 
	@$(CC) $(CFLAGS) -o $@ -c $^

2021/%/day%: parse.o 2021/%/day%.o
	@$(CC) $(CFLAGS) -o $@ $^


clean:
	@rm -rf main *.o