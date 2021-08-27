package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"test/internal/array"
	myUser "test/internal/user"
	myError "test/pkg/error"
	myPostgres "test/pkg/postgres"
)

func main() {
	ctx := context.Background()

	env := os.Getenv("ENVIROMENT")
	if env == "" {
		env = "dev"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("json")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	host := viper.GetString("postgres.databases.jikkosoft.host")
	port := viper.GetInt("postgres.databases.jikkosoft.port")
	user := viper.GetString("postgres.databases.jikkosoft.user")
	passwd := viper.GetString("postgres.databases.jikkosoft.password")
	db := viper.GetString("postgres.databases.jikkosoft.db")
	uri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, passwd, db)

	postgres, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := postgres.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = postgres.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	postgresDriver := myPostgres.PostgresDriver{
		Client: postgres,
	}

	// User
	repoUser := myUser.Repository{
		Finder: &postgresDriver,
	}
	controllerUser := myUser.Controller{
		Context:    ctx,
		Repository: repoUser,
	}

	// Array
	repoArray := array.Repository{
		Sorted: array.Sorted{},
	}
	controllerArray := array.Controller{
		Repository: repoArray,
	}

	// Routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/array", controllerArray.OrderSorted)
	r.Get("/user", controllerUser.GetAllUser)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(myError.Error{Message: "Ruta no encontrada"})
	})

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
