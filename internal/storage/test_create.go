package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) CreateTest(ctx context.Context, text string) error {
	query, params, err := sq.Insert("test").
		Columns(
			"text",
		).
		Values(text).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}

	return nil
}
