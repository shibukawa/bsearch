package bsearch

type CompareResult int

const (
	Smaller CompareResult = -1
	Equal   CompareResult = 0
	Bigger  CompareResult = 1
)

type CompareFunc func(int) CompareResult

func Search(n int, compare CompareFunc) int {
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2
		result := int(compare(h))
		if result < 0 {
			i = h + 1
		} else if result == 0 {
			return h
		} else {
			j = h
		}
	}
	return -1
}

func FindInsertPosition(n int, compare CompareFunc) int {
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2
		result := int(compare(h))
		if result < 0 {
			i = h + 1
		} else if result == 0 {
			return h
		} else {
			j = h
		}
	}
	return i
}
