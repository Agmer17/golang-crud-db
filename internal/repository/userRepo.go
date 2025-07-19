package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Agmer17/golang-crud-db.git/configs"
	"github.com/Agmer17/golang-crud-db.git/internal/model"
)

type UserRepo struct {
	DB *sql.DB
}

func configureDataSource(db *sql.DB, conf configs.AppConfiguration) *sql.DB {
	db.SetMaxOpenConns(conf.MaxConnection)
	db.SetMaxIdleConns(conf.MinIddleConn)
	db.SetConnMaxIdleTime(time.Duration(conf.IdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(conf.LifeTime) * time.Minute)

	return db
}

func NewUserRepo(dbObj *sql.DB, conf configs.AppConfiguration) *UserRepo {
	db := configureDataSource(dbObj, conf)

	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) GetAllData(ctx context.Context) ([]model.UserModel, error) {
	sqlString := "SELECT * FROM PERSON"
	rows, err := repo.DB.QueryContext(ctx, sqlString)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	var listUser []model.UserModel

	for rows.Next() {
		var user model.UserModel

		err := rows.Scan(&user.Id, &user.Nama, &user.Umur, &user.Gender)

		if err != nil {
			return nil, err
		}

		listUser = append(listUser, user)
	}

	return listUser, nil
}

func (repo *UserRepo) AddNewData(newUser model.UserModel, ctx context.Context) (int, error) {
	sqlString := `INSERT INTO person (nama, umur, gender)
					values (?, ?, ?);`
	stmt, err := repo.DB.PrepareContext(ctx, sqlString)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, newUser.Nama, newUser.Umur, newUser.Gender)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	lastID, _ := res.LastInsertId()

	return int(lastID), nil
}

func (repo *UserRepo) DeleteData(ctx context.Context, name string) {

}
