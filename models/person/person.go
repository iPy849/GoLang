package person

import (
	"context"
	"time"

	"fiber.example/database"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type Person struct {
	Id             int64     `json:"Id" `
	FirstName      string    `json:"FirstName" form:"firstName"`
	LastName       string    `json:"LastName" form:"lastName"`
	Email          string    `json:"Email" form:"email"`
	Gender         string    `json:"Gender" form:"gender"`
	DateOfBirth    time.Time `json:"DateOfBirth" form:"dateOfBirth"`
	CountryOfBirth string    `json:"CountryOfBirth" form:"countryOfBirth"`
}

func (p Person) toSlice() []interface{} {
	return []interface{}{p.Id, p.FirstName, p.LastName, p.Email, p.Gender, p.DateOfBirth, p.CountryOfBirth}
}

func RetrieveNPeople(limit, offset int) ([]*Person, error) {
	result := []*Person{}
	query, _, err := database.Pgsq.Select("*").
		From("person").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, err
	}
	queryResult, err := database.DbConnectionPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer queryResult.Close()

	for queryResult.Next() {
		// TODO: Aquí no jala, por qué???
		p := new(Person)
		queryResult.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Email, &p.Gender, &p.DateOfBirth, &p.CountryOfBirth)
		result = append(result, p)
	}
	return result, nil
}

func RetrievePerson(id int) (*Person, error) {
	query, _, err := database.Pgsq.Select("*").
		From("person").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}
	queryResult := database.DbConnectionPool.QueryRow(context.Background(), query, id)
	p := new(Person)
	err = queryResult.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Email, &p.Gender, &p.DateOfBirth, &p.CountryOfBirth)

	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}
	return p, nil
}

func RetrievePersonByNames(firstName, lastName string) (*Person, error) {
	query, _, err := database.Pgsq.
		Select("id", "first_name", "last_name").
		From("person").
		Where(
			sq.Eq{
				"first_name": firstName,
				"last_name":  lastName,
			},
		).
		ToSql()
	queryResult := database.DbConnectionPool.QueryRow(context.Background(), query, firstName, lastName)
	p := new(Person)
	err = queryResult.Scan(&p.Id, &p.FirstName, &p.LastName)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return p, nil
}

// TODO: Biometría Hemática

func CreatePerson(p *Person) (int64, error) {
	insertionData := p.toSlice()[1:]
	query, _, err := database.Pgsq.Insert("person").
		Columns("first_name", "last_name", "email", "gender", "date_of_birth", "country_of_birth").
		Values(insertionData...).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}
	queryResult := database.DbConnectionPool.QueryRow(context.Background(), query, insertionData...)
	var insertId int64
	err = queryResult.Scan(&insertId)
	if err != nil {
		return 0, err
	}
	return insertId, nil
}

func UpdatePerson(id int, p *Person) error {
	emptyPerson := Person{}
	updateData := []interface{}{}
	queryBuilder := database.Pgsq.Update("person")

	if p.FirstName != emptyPerson.FirstName {
		queryBuilder = queryBuilder.Set("first_name", p.FirstName)
		updateData = append(updateData, p.FirstName)
	}
	if p.LastName != emptyPerson.LastName {
		queryBuilder = queryBuilder.Set("last_name", p.LastName)
		updateData = append(updateData, p.LastName)
	}
	if p.Email != emptyPerson.Email {
		queryBuilder = queryBuilder.Set("email", p.Email)
		updateData = append(updateData, p.Email)
	}
	if p.Gender != emptyPerson.Gender {
		queryBuilder = queryBuilder.Set("gender", p.Gender)
		updateData = append(updateData, p.Gender)
	}
	if p.DateOfBirth != emptyPerson.DateOfBirth {
		queryBuilder = queryBuilder.Set("date_of_birth", p.DateOfBirth)
		updateData = append(updateData, p.DateOfBirth)
	}
	if p.CountryOfBirth != emptyPerson.CountryOfBirth {
		queryBuilder = queryBuilder.Set("country_of_birth", p.CountryOfBirth)
		updateData = append(updateData, p.CountryOfBirth)
	}

	query, _, err := queryBuilder.
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}
	updateData = append(updateData, id)
	_, err = database.DbConnectionPool.Exec(context.Background(), query, updateData...)
	if err != nil {
		return err
	}
	return nil
}

func DeletePerson(id int) error {
	query, _, err := database.Pgsq.Delete("person").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}
	_, err = database.DbConnectionPool.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
