package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/vinbyte/protip-golang-testing/domain"
)

type postgresRepository struct {
	conn *sql.DB
}

// NewUserRepository is used to init new user repository
func NewUserRepository(pgConn *sql.DB) domain.UserRepository {
	return &postgresRepository{
		conn: pgConn,
	}
}

func (p *postgresRepository) Fetch(ctx context.Context, userID int) ([]domain.User, error) {
	var result []domain.User
	var err error

	query := "SELECT * from users"
	if userID > 0 {
		query += fmt.Sprintf(" WHERE id = %d", userID)
	}
	rows, err := p.conn.QueryContext(ctx, query)
	if err != nil {
		return result, err
	}
	defer rows.Close()
	for rows.Next() {
		temp := domain.User{}
		err := rows.Scan(&temp.ID, &temp.Name, &temp.Email)
		if err != nil {
			return result, err
		}
		result = append(result, temp)
	}

	return result, err
}
