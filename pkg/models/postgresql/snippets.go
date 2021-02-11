package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"snippetBoxReborn/pkg/models"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES($1, $2, now(), now() + make_interval(days => $3)) RETURNING id;"
	id := 0

	err := m.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := "SELECT id, title, content, created, expires FROM snippets WHERE AND id = $1;"
	row := m.DB.QueryRow(context.Background(), stmt, id)
	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
