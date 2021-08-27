package user

import (
	"context"
	"encoding/json"
	"net/http"
	myError "test/pkg/error"
)

type (
	Controller struct {
		Context    context.Context
		Repository Repository
	}
)

func (controller Controller) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	array, err := controller.Repository.GetAllUsers(controller.Context)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(myError.Error{
			Message: err.Error(),
		})
	} else {
		json.NewEncoder(w).Encode(array)
	}
}
