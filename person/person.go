package person

import (
	"context"
	"time"

	"fiber.example/database"
)

type Person struct {
	FirstName      string
	LastName       string
	Email          string
	Gender         string
	DateOfBirth    time.Time
	CountryOfBirth string
}

func RetrieveNPeople(limit, offset int) ([]*Person, error) {
	result := []*Person{}

	statement := "SELECT * FROM person LIMIT $1 OFFSET $2"
	queryResult, err := database.DbConnectionPool.Query(context.Background(), statement, limit, offset)
	if err != nil {
		return nil, err
	}
	defer queryResult.Close()

	for queryResult.Next() {
		vals, err := queryResult.Values()
		if err != nil {
			return nil, err
		}
		p := Person{
			FirstName:      vals[1].(string),
			LastName:       vals[2].(string),
			Email:          vals[3].(string),
			Gender:         vals[4].(string),
			DateOfBirth:    vals[5].(time.Time),
			CountryOfBirth: vals[6].(string),
		}
		result = append(result, &p)
	}
	return result, nil
}

func RetrievePeople(id int) (*Person, error) {
	statement := "SELECT * FROM person WHERE id=$1;"
	queryResult := database.DbConnectionPool.QueryRow(context.Background(), statement, id)
	p := Person{}
	err := queryResult.Scan(&id, &p.FirstName, &p.LastName, &p.Email, &p.Gender, &p.DateOfBirth, &p.CountryOfBirth)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
