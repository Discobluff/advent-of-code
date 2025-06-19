CC = clang
CFLAGS = -Wall -Wextra -Wpedantic -std=c99 #-fsanitize=address -g

../utils/%.o: ../utils/%.c
	@$(CC) $(CFLAGS) -c $^ -o $@

%/day.o: %/day.c 
	@$(CC) $(CFLAGS) -o $@ -c $^

%/main.o: %/main.c
	@$(CC) $(CFLAGS) -o $@ -c $^

%/main: ../utils/parse.o %/main.o %/day.o
	@$(CC) $(CFLAGS) -o $@ $^

%/test.o: %/test.c
	@$(CC) $(CFLAGS) -o $@ -c $^

%/test: ../utils/parse.o %/test.o %/day.o
	@$(CC) $(CFLAGS) -o $@ $^ -lcunit

.SILENT:
run%: %/main
	@./$^

test%: %/test
	@./$^

clean:
	@rm -rf main *.o