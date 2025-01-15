package priorityQueue

import "fmt"

type pQueue[T any] []T

func DefQueue[T any]() pQueue[T] {
	var res pQueue[T]
	return res
}

func IsEmptyQueue[T any](queue pQueue[T]) bool {
	return len(queue) == 0
}

func AddQueue[T any](q *pQueue[T], elem T, cmp func(T, T) bool) {
	var index int
	for index < len(*q) && cmp((*q)[index], elem) {
		index++
	}
	if index >= len(*q) {
		(*q) = append((*q), elem)
	} else {
		(*q) = append((*q), (*q)[len(*q)-1])
		for i := len(*q) - 2; i > index; i-- {
			(*q)[i] = (*q)[i-1]
		}
		(*q)[index] = elem
	}
}

func PopQueue[T any](q *pQueue[T]) T {
	if len(*q) == 0 {
		panic("cannot pop from an empty queue")
	}
	var res T = (*q)[0]
	(*q) = (*q)[1:]
	return res
}

func Display[T any](q pQueue[T]) {
	for _, val := range q {
		fmt.Println(val)
	}
}
