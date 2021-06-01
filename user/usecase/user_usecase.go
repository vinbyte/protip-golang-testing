package usecase

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/vinbyte/protip-golang-testing/app/helpers"
	"github.com/vinbyte/protip-golang-testing/domain"
)

type userUsecase struct {
	repo           domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(r domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		repo:           r,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetAll(ctx context.Context) (code int, response interface{}) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var resp helpers.BaseResponse
	resp.Code = http.StatusOK
	resp.Message = "success"
	resp.Data = new(struct{})
	var userList domain.UserList
	userList.Items = make([]domain.User, 0)
	defer func() {
		code = resp.Code
		response = resp
	}()
	usersData, err := u.repo.Fetch(ctx, 0)
	if err != nil {
		log.Println(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		return
	}
	if len(usersData) == 0 {
		resp.Code = 404
		resp.Message = "not found"
		return
	}

	userList.Items = usersData
	resp.Data = userList

	return
}

func (u *userUsecase) GetByID(ctx context.Context, userID int) (code int, response interface{}) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var resp helpers.BaseResponse
	resp.Code = http.StatusOK
	resp.Message = "success"
	resp.Data = new(struct{})
	defer func() {
		code = resp.Code
		response = resp
	}()

	if userID == 0 {
		resp.Code = 404
		resp.Message = "not found"
		return
	}
	usersData, err := u.repo.Fetch(ctx, userID)
	if err != nil {
		log.Println(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = err.Error()
		return
	}
	if len(usersData) == 0 {
		resp.Code = 404
		resp.Message = "not found"
		return
	}
	resp.Data = usersData[0]
	return
}
