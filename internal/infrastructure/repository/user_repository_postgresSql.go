package repository

import (
	"go-crud/internal/domain/entity"

	"database/sql"

	_ "github.com/lib/pq"
)

type UserPostgress struct {
	Client *sql.DB
}

func NewUserPostgress(client *sql.DB) *UserPostgress {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")

	if err != nil {
		panic(err)
	}

	userDB := &UserPostgress{
		Client: client,
	}

	return userDB
}

func (db *UserPostgress) Insert(u entity.User) (id string, err error) {
	conn := db.Client
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO categorias (nome) VALUES ($1) RETURNING id`

	err = conn.QueryRow(sql, u.Name).Scan(&id)

	return id, nil
}

func (db *UserPostgress) Get(id string) (u entity.User, err error) {
	conn := db.Client
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM users WHERE id=$1`, id)

	err = row.Scan(&u.ID, &u.Name)

	return 
}

func (db *UserPostgress) GetAll() (sc []entity.User, err error) {
	conn := db.Client

	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM categorias`)
	if err != nil {
		return
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

	return
}

func (db *UserPostgress) Update(id string, u entity.User) (*entity.User, error) {
	conn := db.Client
	defer conn.Close()

	_ , err := conn.Exec(`UPDATE users SET nome=$2 WHERE id=$1`, id, u.Name)
	if err != nil {
		return nil, err
	}

	new, err := db.Get(id)
	if err != nil {
		return nil, err
	}

	return &new, err
}

func (db *UserPostgress) Delete(id int64) (int64, error) {
	conn := db.Client
	if err != nil {
		return
	}
	defer conn.Close()

	res, err := db.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return res.RowsAffected()
}
