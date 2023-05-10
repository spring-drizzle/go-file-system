package config

import "fmt"

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	pass = "postgres"
	db   = "postgres"
)

func GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, db)
}
