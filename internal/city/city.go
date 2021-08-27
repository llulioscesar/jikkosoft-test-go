package city

import "github.com/google/uuid"

type (
	City struct {
		ID   uuid.UUID `json:"id" pg:"type:uuid,default:uuid_generate_v4(),column_name:id"`
		Name string    `json:"name" pg:"varchar(55),column_name:name"`
	}
)
