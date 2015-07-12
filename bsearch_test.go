package bsearch

import (
	"testing"
)

func Test_Search_Found(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 11
	result := Search(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != 5 {
		t.Error("found position was wrong:", result)
	}
}

func Test_Search_NotFound_1(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 12
	result := Search(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != -1 {
		t.Error("found position was wrong:", result)
	}
}

func Test_Search_NotFound_2(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 0
	result := Search(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != -1 {
		t.Error("found position was wrong:", result)
	}
}

func Test_Search_NotFound_3(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 24
	result := Search(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != -1 {
		t.Error("found position was wrong:", result)
	}
}

func Test_FindInsertPosition_Found(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 11
	result := FindInsertPosition(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != 5 {
		t.Error("found position was wrong:", result)
	}
}

func Test_FindInsertPosition_NotFound_1(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 12
	result := FindInsertPosition(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != 6 {
		t.Error("found position was wrong:", result)
	}
}

func Test_FindInsertPosition_NotFound_2(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 0
	result := FindInsertPosition(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != 0 {
		t.Error("found position was wrong:", result)
	}
}

func Test_FindInsertPosition_NotFound_3(t *testing.T) {
	items := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	target := 24
	result := FindInsertPosition(len(items), func(i int) CompareResult {
		if target > items[i] {
			return Smaller
		} else if target == items[i] {
			return Equal
		} else {
			return Bigger
		}
	})
	if result != 10 {
		t.Error("found position was wrong:", result)
	}
}