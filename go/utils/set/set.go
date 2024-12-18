package set

type set[T comparable] map[T]struct{}

func DefSet[T comparable]() set[T] {
	return make(map[T]struct{})
}

func In[T comparable](s set[T], elem T) bool {
	var _, ok = s[elem]
	return ok
}

func Add[T comparable](s set[T], elem T) {
	s[elem] = struct{}{}
}

func Remove[T comparable](s set[T], elem T) {
	delete(s, elem)
}

func Union[T comparable](s1 set[T], s2 set[T]) set[T] {
	var s set[T] = DefSet[T]()
	for key := range s1 {
		Add(s, key)
	}
	for key := range s2 {
		Add(s, key)
	}
	return s
}

func Intersect[T comparable](s1 set[T], s2 set[T]) set[T] {
	var res set[T] = DefSet[T]()
	for key := range s1 {
		if In(s2, key) {
			Add(res, key)
		}
	}
	return res
}

func Include[T comparable](s1 set[T], s2 set[T]) bool {
	for key := range s1 {
		if !In(s2, key) {
			return false
		}
	}
	return true
}

func Deprived[T comparable](s1 set[T], s2 set[T]) set[T] {
	var res set[T] = DefSet[T]()
	for key := range s1 {
		if !In(s2, key) {
			Add(res, key)
		}
	}
	return res
}
