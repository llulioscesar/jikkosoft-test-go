package user

import (
	"github.com/google/uuid"
	"test/internal/city"
)

type (
	User struct {
		ID       uuid.UUID `json:"id" pg:"type:uuid,default:uuid_generate_v4(),column_name:id"`
		Name     string    `json:"name" pg:"varchar(55),column_name:name"`
		LastName string    `json:"last_name" pg:"varchar(55),column_name:last_name"`
		Years    int       `json:"years" pg:"int,column_name:years"`
		Gender   string    `json:"gender" pg:"varchar(55),column_name:gender"`
		City     city.City `json:"city"`
	}
)
