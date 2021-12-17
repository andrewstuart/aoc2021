package myaoc

func GroupByFunc[T, V any, K comparable](ts []T, f func(T) (K, V)) map[K][]V {
	m := map[K][]V{}
	for _, t := range ts {
		k, v := f(t)
		m[k] = append(m[k], v)
	}
	return m
}

func MapOfFunc[T, V any, K comparable](ts []T, f func(T) (K, V)) map[K]V {
	m := map[K]V{}
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}
