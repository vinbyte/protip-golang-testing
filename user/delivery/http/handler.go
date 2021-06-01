package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/vinbyte/protip-golang-testing/app/helpers"
	"github.com/vinbyte/protip-golang-testing/domain"
)

// UserHandler we export this struct for make testing easier
type UserHandler struct {
	UserUsecase domain.UserUsecase
}

// NewUserHandler used for init new handler
func NewUserHandler(c *chi.Mux, uu domain.UserUsecase) {
	handler := UserHandler{
		UserUsecase: uu,
	}
	c.Route("/v1", func(r chi.Router) {
		r.Get("/users", handler.UserList)
		r.Get("/users/{id}", handler.UserByID)
	})
}

// UserList handle the user list request
func (u *UserHandler) UserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code, response := u.UserUsecase.GetAll(ctx)
	helpers.SendResponse(w, code, response)
}

// UserByID get the user by it's id
func (u *UserHandler) UserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		userID = 0
	}
	code, response := u.UserUsecase.GetByID(ctx, userID)
	helpers.SendResponse(w, code, response)
}
