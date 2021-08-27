package array

import (
	"errors"
	"sort"
)

type (
	Sorted struct {
	}
)

func (s Sorted) Process(array []int) ([]int, error) {
	sort.Ints(array)

	size := len(array)

	var numbas []int
	if size > 0 {
		numbas = append(numbas, array[0])
		var duplas []int
		for i, s := range array {
			next := i + 1
			if next >= size {
				next = i
			}
			equal := s == array[next]
			if i != next {
				if !equal {
					numbas = append(numbas, array[next])
				} else {
					duplas = append(duplas, array[next])
				}
			}
		}
		numbas = append(numbas, duplas...)
		return numbas, nil
	}
	return nil, errors.New("Array unsorted is empty.")
}
