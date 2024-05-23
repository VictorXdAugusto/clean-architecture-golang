package repository

import (
	"database/sql"
	"go-crud/internal/domain/entity"
)

type UserPostgress struct {
	Client *sql.DB
}

func NewUserPostgress(client *sql.DB) *UserPostgress {
	return &UserPostgress{
		Client: client,
	}
}

func (db *UserPostgress) Insert(u entity.User) (id string, err error) {
	sql := `INSERT INTO "user" (nome) VALUES ($1) RETURNING id`
	err = db.Client.QueryRow(sql, u.Name).Scan(&id)
	return id, err
}

func (db *UserPostgress) Get(id string) (u entity.User, err error) {
	row := db.Client.QueryRow(`SELECT * FROM "user" WHERE id=$1`, id)
	err = row.Scan(&u.ID, &u.Name)
	return u, err
}

func (db *UserPostgress) GetAll() (sc []entity.User, err error) {
	rows, err := db.Client.Query(`SELECT * FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u entity.User
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			continue
		}
		sc = append(sc, u)
	}

	return sc, rows.Err()
}

func (db *UserPostgress) Update(id string, u entity.User) (*entity.User, error) {
	_, err := db.Client.Exec(`UPDATE "user" SET nome=$2 WHERE id=$1`, id, u.Name)
	if err != nil {
		return nil, err
	}

	updatedUser, err := db.Get(id)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (db *UserPostgress) Delete(id string) (int64, error) { // changed id type to string for consistency
	res, err := db.Client.Exec(`DELETE FROM "user" WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
