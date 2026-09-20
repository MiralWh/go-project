package helperstructs

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func Create[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		data: make(map[T]struct{}),
	}
}

func (s *HashSet[T]) Insert(value T) {
	s.data[value] = struct{}{}
}

func (s *HashSet[T]) Remove(value T) {
	delete(s.data, value)
}

func (s *HashSet[T]) Exists(value T) bool {
	_, exists := s.data[value]
	return exists
}

func (s *HashSet[T]) Length() int {
	return len(s.data)
}

func (s *HashSet[T]) Sliced() []T {
	slice := make([]T, 0, len(s.data))
	for v := range s.data {
		slice = append(slice, v)
	}
	return slice
}
