package gochapter

func addAll[T any](source *[]T, elements []T) {
	for _, el := range elements {
		*source = append(*source, el)
	}
}
