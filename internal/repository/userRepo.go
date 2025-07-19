package repository

import (
	"context"
	"database/sql"
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
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

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
