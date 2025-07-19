package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Agmer17/golang-crud-db.git/internal/model"
	"github.com/Agmer17/golang-crud-db.git/internal/service"
	"github.com/Agmer17/golang-crud-db.git/internal/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Error model.ErrorResponse
type Success model.SuccessResponse

type UserController struct {
	svc *service.UserService
}

func (controller *UserController) GetAllData(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)

	defer cancel()

	data, err := controller.svc.GetAllData(ctx)

	if err != nil {
		ErrorResp := Error{
			Status:     "INTERNAL SERVER ERROR",
			Detail:     "TERJADI KESALAHAN SAAT MELAKUKAN OPERASI KE DATABASE. HARAP CEK KONFIGURASI DAN QUERY YANG DIJALANAKAN",
			Errors:     err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
		log.Fatal(err)
		util.WriteJson(w, ErrorResp, ErrorResp.StatusCode)
	}

	SucessResp := Success{
		Status:     "OK",
		Detail:     "Berhasil mengambil data dari database",
		Data:       data,
		StatusCode: http.StatusOK,
	}
	util.WriteJson(w, SucessResp, SucessResp.StatusCode)

}

func NewUserController(svc *service.UserService) *UserController {
	return &UserController{
		svc: svc,
	}
}

func (ctrl *UserController) RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/user", func(r chi.Router) {
		r.Get("/get-all", ctrl.GetAllData)
	})
	return r
}
