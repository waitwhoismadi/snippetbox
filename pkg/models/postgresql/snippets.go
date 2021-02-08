package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"snippetBoxReborn/pkg/models"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	id := 0
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES($1, $2, now(), now() + INTERVAL '$3' RETURNING id);"

	err := m.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
