package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
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

	_, err = postgres.ExecContext(ctx, `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS city (
		    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		    name VARCHAR(55)
		);
		CREATE TABLE IF NOT EXISTS "user" (
		    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
		    name VARCHAR(55),
		    last_name VARCHAR(55),
		    year INT,
		    Gender VARCHAR(55),
		    city_id UUID NOT NULL,
		    FOREIGN KEY(city_id) REFERENCES city(id)
		);
	`)

	if err != nil{
		panic(err)
	}

	fmt.Println("Base de datos creada")

	_, err = postgres.ExecContext(ctx, `
		INSERT INTO city (id, name) VALUES 
		   ('0563825c-8982-408f-b7a6-14180b5dc7d0', 'BOGOTA'),
		   ('39084526-a04d-4b81-9f96-7d0dae728822', 'CALI'),
		   ('fc1f1c22-76cf-402a-8d9f-7200da9b7f36', 'MEDELLIN');

		INSERT INTO "user" (name, last_name, year, gender, city_id) VALUES 
			('LUISITO', 'COMUNICA', 25, 'MASCULINO', '0563825c-8982-408f-b7a6-14180b5dc7d0'),
		    ('MARIA', 'CASTRO', 35, 'FEMENINO', '39084526-a04d-4b81-9f96-7d0dae728822'),
			('JOSE', 'CACERES', 28, 'MASCULINO', 'fc1f1c22-76cf-402a-8d9f-7200da9b7f36');
	`)
	if err != nil{
		panic(err)
	}
	fmt.Println("Datos creados")

}
