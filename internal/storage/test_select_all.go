package storage

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) GetTests(ctx context.Context) ([]Test, error) {
	query, params, err := sq.Select(
		"id",
		"text",
	).From("test").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var dest []Test
	err = s.db.SelectContext(ctx, &dest, query, params...)
	if err != nil {
		return nil, err
	}

	return dest, err
}
