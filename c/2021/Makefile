CC = clang
CFLAGS = -Wall -Wextra -Wpedantic -std=c99 -fsanitize=address -g
LIBSTEST = -L /cunit/CUnit -lcunit
OLIBSTEST = -I /cunit/CUnit

../utils/%.o: ../utils/%.c
	@$(CC) $(CFLAGS) -c $^ -o $@

%/day.o: %/day.c 
	@$(CC) $(CFLAGS) -o $@ -c $^

%/main.o: %/main.c
	@$(CC) $(CFLAGS) -o $@ -c $^

%/main: ../utils/parse.o %/main.o %/day.o
	@$(CC) $(CFLAGS) -o $@ $^

%/test.o: %/test.c
	@$(CC) $(CFLAGS) -o $@ -c $^ $(OLIBSTEST)

%/test: ../utils/parse.o %/test.o %/day.o
	@$(CC) $(CFLAGS) -o $@ $^ -lcunit $(LIBSTEST)

.SILENT:
run%: %/main
	@./$^

test%: %/test
	@./$^

clean:
	@rm -rf main *.o

MAKEFLAGS += --no-print-directory

run:
	for dir in $(shell find . -type d); do \
	if [ "$$dir" = "." ]; then continue; fi; \
		day=$$(basename $$dir); \
		echo "2021 Day $$day"; \
		make run$$day; \
		echo ""; \
	done

test:
	for dir in $(shell find . -type d); do \
	if [ "$$dir" = "." ]; then continue; fi; \
		day=$$(basename $$dir); \
		echo "2021 Day $$day"; \
		make test$$day; \
		echo ""; \
	done

