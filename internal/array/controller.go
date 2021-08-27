package array

import (
	"encoding/json"
	"net/http"
	myError "test/pkg/error"
)

type (
	Controller struct {
		Repository Repository
	}
)

func (controller Controller) OrderSorted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var body ArrayRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(myError.Error{
			Message: err.Error(),
		})
	} else {
		array, err := controller.Repository.Shorted(body.Unsorted)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(myError.Error{
				Message: err.Error(),
			})
		} else {
			response := ArrayResponse{
				Unsorted: body.Unsorted,
				Sorted:   array,
			}
			json.NewEncoder(w).Encode(response)
		}
	}

}
