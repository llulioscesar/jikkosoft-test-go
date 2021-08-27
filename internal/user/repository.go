package user

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	city2 "test/internal/city"
)

type (
	Repository struct {
		Finder interface {
			Query(context.Context, string) (*sql.Rows, error)
		}
	}
)

func (repo Repository) GetAllUsers(ctx context.Context) ([]User, error) {
	var users []User
	rows, err := repo.Finder.Query(ctx, `
		select "user".id,
			   "user".name,
               
			   "user".year,
			   "user".gender,
			   city.id,
			   city.name
		from "user"
				 inner join city city on city.id = "user".city_id
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var error error

	for rows.Next() {
		var userId string
		var name string
		var lastName string
		var year int
		var gender string
		var cityName string
		var cityId string

		err := rows.Scan(&userId, &name, &lastName, &year, &gender, &cityId, &cityName)
		if err != nil {
			error = err
		} else {
			city := city2.City{
				ID:   uuid.MustParse(cityId),
				Name: cityName,
			}
			user := User{
				ID:       uuid.MustParse(userId),
				Name:     name,
				LastName: lastName,
				Years:    year,
				Gender:   gender,
				City:     city,
			}
			users = append(users, user)
		}
	}
	if error != nil {
		return nil, error
	}
	return users, nil
}
