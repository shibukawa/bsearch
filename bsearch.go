package bsearch

type CompareResult int

const (
	Smaller CompareResult = -1
	Equal CompareResult = 0
	Bigger CompareResult = 1
)

type CompareFunc func(int)CompareResult

func Search(n int, compare CompareFunc) int {
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2
		switch compare(h) {
		case Smaller:
			i = h + 1
		case Equal:
			return h
		case Bigger:
			j = h
		}
	}
	return -1
}

func FindInsertPosition(n int, compare CompareFunc) int {
	i, j := 0, n
	for i < j {
		h := i + (j-i)/2
		switch compare(h) {
		case Smaller:
			i = h + 1
		case Equal:
			return h
		case Bigger:
			j = h
		}
	}
	return i
}