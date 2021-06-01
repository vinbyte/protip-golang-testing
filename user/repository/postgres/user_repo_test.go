package postgres_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinbyte/protip-golang-testing/domain"
	"github.com/vinbyte/protip-golang-testing/user/repository/postgres"
)

func TestFetch(t *testing.T) {
	t.Run("success fetch all", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		assert.Nil(t, err)
		query := "SELECT * from users"
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
		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(mockUsers[0].ID, mockUsers[0].Name, mockUsers[0].Email).
			AddRow(mockUsers[1].ID, mockUsers[1].Name, mockUsers[1].Email)
		mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

		repo := postgres.NewUserRepository(mockDB)
		results, err := repo.Fetch(context.Background(), 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(results))
	})
}
