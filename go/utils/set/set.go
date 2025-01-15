package set

import "maps"

type Set[T comparable] map[T]struct{}

func DefSet[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func In[T comparable](s Set[T], elem T) bool {
	var _, ok = s[elem]
	return ok
}

func Add[T comparable](s Set[T], elem T) {
	s[elem] = struct{}{}
}

func Remove[T comparable](s Set[T], elem T) {
	delete(s, elem)
}

func Equal[T comparable](s1 Set[T], s2 Set[T]) bool {
	return maps.Equal(s1, s2)
}

func Without[T comparable](s Set[T], e T) Set[T] {
	var newS = DefSet[T]()
	Add(newS, e)
	return Deprived(s, newS)
}

func Max[T comparable](s Set[T], f func(T, T) bool) (T, bool) {
	var res T
	if IsEmpty(s) {
		return res, false
	}
	var found bool
	for elem := range s {
		if !found || f(res, elem) {
			found = true
			res = elem
		}
	}
	return res, true
}

func SetToSlice[T comparable](s Set[T]) []T {
	var res []T = make([]T, len(s))
	var index int
	for e := range s {
		res[index] = e
		index++
	}
	return res
}

func Union[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	var s Set[T] = DefSet[T]()
	for key := range s1 {
		Add(s, key)
	}
	for key := range s2 {
		Add(s, key)
	}
	return s
}

func Intersect[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	var res Set[T] = DefSet[T]()
	for key := range s1 {
		if In(s2, key) {
			Add(res, key)
		}
	}
	return res
}

func Include[T comparable](s1 Set[T], s2 Set[T]) bool {
	for key := range s1 {
		if !In(s2, key) {
			return false
		}
	}
	return true
}

func Deprived[T comparable](s1 Set[T], s2 Set[T]) Set[T] {
	var res Set[T] = DefSet[T]()
	for key := range s1 {
		if !In(s2, key) {
			Add(res, key)
		}
	}
	return res
}

func IsEmpty[T comparable](s Set[T]) bool {
	return len(s) == 0
}
