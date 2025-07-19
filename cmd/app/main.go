package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Agmer17/golang-crud-db.git/configs"
	"github.com/Agmer17/golang-crud-db.git/internal/controller"
	"github.com/Agmer17/golang-crud-db.git/internal/repository"
	"github.com/Agmer17/golang-crud-db.git/internal/service"
)

func main() {
	configs.LoadEnv()
	appConfig := configs.NewConfig()
	db, err := sql.Open("mysql", appConfig.DbUrl)
	if err != nil {
		panic(err)
	}
	userRepo := repository.NewUserRepo(db, *appConfig)
	userService := service.NewUserService(userRepo)
	crl := controller.NewUserController(userService)
	fmt.Println("aplikasi berjalan di port", appConfig.ServerLocation)
	log.Fatal(http.ListenAndServe(appConfig.ServerLocation, crl.RegisterRoutes()))

}
