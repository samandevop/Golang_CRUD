package storage

import (
	"database/sql"
	"fmt"

	"crud/models"
)

func Create(db *sql.DB, user models.User) (string, error) {

	var (
		id    string
		query string
	)

	query = `
		INSERT INTO 
			users (first_name, last_name)
		VALUES ( $1, $2 )
		RETURNING id
	`
	err := db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetById(db *sql.DB, id string) (models.User, error) {

	var (
		user  models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
		WHERE id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetList(db *sql.DB) ([]models.User, error) {

	var (
		users []models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.User{}, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func Update(db *sql.DB, user models.User) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE
			users
		SET
			first_name = $2,
			last_name = $3
		WHERE
			id = $1
	`

	result, err := db.Exec(
		query,
		user.Id,
		user.FirstName,
		user.LastName,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func Patch(db *sql.DB, user models.UserData) (int64, error) {

	var (
		query = ""
		set   = ""
	)

	query = `
		UPDATE
			users
		SET `

	for key, val := range user.Data {

		switch val.(type) {
		case int:
			set += fmt.Sprintf("%s = %d,", key, val)
		case string:
			set += fmt.Sprintf("%s = '%s',", key, val)
		case bool:
			set += fmt.Sprintf("%s = %v,", key, val)
		case float64:
			set += fmt.Sprintf("%s = %v,", key, val)
		}
	}

	query += set

	query = query[:len(query)-1]

	query += " WHERE id = $1"

	fmt.Println(query)

	result, err := db.Exec(query, user.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func Delete(db *sql.DB, id string) error {

	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
