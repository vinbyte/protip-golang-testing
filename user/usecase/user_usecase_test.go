package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vinbyte/protip-golang-testing/app/helpers"
	"github.com/vinbyte/protip-golang-testing/domain"
	"github.com/vinbyte/protip-golang-testing/domain/mocks"
	"github.com/vinbyte/protip-golang-testing/user/usecase"
)

func TestGetAll(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	ctx := context.Background()

	mockUsers := []domain.User{
		{
			ID:    1,
			Name:  "people 1",
			Email: "mail1@com.com",
		},
		{
			ID:    2,
			Name:  "people 2",
			Email: "mail2@com.com",
		},
	}
	usecase := usecase.NewUserUsecase(mockRepo, time.Second*3)
	mockRepo.On("Fetch", mock.Anything, 0).Return(mockUsers, nil)

	code, response := usecase.GetAll(ctx)
	basicResponse := response.(helpers.BaseResponse)
	dataResponse := basicResponse.Data.(domain.UserList)
	assert.Equal(t, 200, code)
	assert.Equal(t, 200, basicResponse.Code)
	assert.Len(t, dataResponse.Items, 2)
}

func TestGetByID(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	ctx := context.Background()

	mockUsers := []domain.User{
		{
			ID:    1,
			Name:  "people 1",
			Email: "mail1@com.com",
		},
	}
	usecase := usecase.NewUserUsecase(mockRepo, time.Second*3)
	mockRepo.On("Fetch", mock.Anything, 1).Return(mockUsers, nil)

	code, response := usecase.GetByID(ctx, 1)
	basicResponse := response.(helpers.BaseResponse)
	dataResponse := basicResponse.Data.(domain.User)
	assert.Equal(t, 200, code)
	assert.Equal(t, 200, basicResponse.Code)
	assert.Equal(t, mockUsers[0].Name, dataResponse.Name)
}
