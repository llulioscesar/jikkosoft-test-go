package array

type (
	ArrayRequest struct {
		Unsorted []int `json:"unsorted"`
	}

	ArrayResponse struct {
		Unsorted []int `json:"unsorted"`
		Sorted   []int `json:"sorted"`
	}
)
