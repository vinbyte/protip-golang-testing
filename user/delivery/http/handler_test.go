package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/vinbyte/protip-golang-testing/app/helpers"
	"github.com/vinbyte/protip-golang-testing/domain"
	"github.com/vinbyte/protip-golang-testing/domain/mocks"
	httpHandler "github.com/vinbyte/protip-golang-testing/user/delivery/http"
)

func TestUserList(t *testing.T) {
	mockUsecase := new(mocks.UserUsecase)
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
	var resp helpers.BaseResponse
	resp.Code = 200
	resp.Message = "success"
	resp.Data = domain.UserList{
		Items: mockUsers,
	}
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/v1/users", nil)
	assert.Nil(t, err)
	mockUsecase.On("GetAll", r.Context()).Return(resp.Code, resp)

	h := httpHandler.UserHandler{
		UserUsecase: mockUsecase,
	}
	h.UserList(w, r)
	var expectedResponse helpers.BaseResponse
	json.NewDecoder(w.Body).Decode(&expectedResponse)
	expectedData := expectedResponse.Data.(map[string]interface{})
	assert.Equal(t, resp.Code, w.Code)
	assert.Equal(t, 200, expectedResponse.Code)
	assert.Equal(t, 2, reflect.ValueOf(expectedData["items"]).Len())
}

func TestUserByID(t *testing.T) {
	mockUsecase := new(mocks.UserUsecase)
	mockUsers := []domain.User{
		{
			ID:    1,
			Name:  "people 1",
			Email: "mail1@com.com",
		},
	}
	var resp helpers.BaseResponse
	resp.Code = 200
	resp.Message = "success"
	resp.Data = mockUsers[0]
	mockUserID := 1
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/v1/users/"+strconv.Itoa(mockUserID), nil)
	assert.Nil(t, err)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", strconv.Itoa(mockUserID))
	ctx := context.WithValue(context.Background(), chi.RouteCtxKey, chiContext)
	r = r.WithContext(ctx)
	mockUsecase.On("GetByID", r.Context(), 1).Return(resp.Code, resp)

	h := httpHandler.UserHandler{
		UserUsecase: mockUsecase,
	}
	h.UserByID(w, r)
	var expectedResponse helpers.BaseResponse
	json.NewDecoder(w.Body).Decode(&expectedResponse)
	expectedData := expectedResponse.Data.(map[string]interface{})
	assert.Equal(t, resp.Code, w.Code)
	assert.Equal(t, 200, expectedResponse.Code)
	assert.Equal(t, mockUsers[0].Name, expectedData["name"])
}
