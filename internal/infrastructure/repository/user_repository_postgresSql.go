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
	sql := `INSERT INTO "users" (id, name, email, username, password, created_at, updated_at, age, is_active) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	err = db.Client.QueryRow(sql, u.ID, u.Name, u.Email, u.Username, u.Password, u.CreatedAt, u.UpdatedAt, u.Age, u.IsActive).Scan(&id)
	return id, err
}

func (db *UserPostgress) Get(id string) (u entity.User, err error) {
	row := db.Client.QueryRow(`SELECT id, name, email, username, password, created_at, updated_at, age, is_active FROM "users" WHERE id=$1`, id)
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.Age, &u.IsActive)
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
	_, err := db.Client.Exec(`
		UPDATE "users" 
		SET name=$2, email=$3, username=$4, password=$5, created_at=$6, updated_at=$7, age=$8, is_active=$9 
		WHERE id=$1
	`, id, u.Name, u.Email, u.Username, u.Password, u.CreatedAt, u.UpdatedAt, u.Age, u.IsActive)
	if err != nil {
		return nil, err
	}

	updatedUser, err := db.Get(id)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (db *UserPostgress) Delete(id string) (int64, error) { 
	res, err := db.Client.Exec(`DELETE FROM "users" WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
