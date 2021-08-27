package array

type (
	Repository struct {
		Sorted Sorted
	}
)

func (repo Repository) Shorted(array []int) ([]int, error) {
	return repo.Sorted.Process(array)
}
