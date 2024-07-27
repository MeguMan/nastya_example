package storage

import (
	"context"
)

func (s *Storage) GetTestByID(ctx context.Context, id int) (Test, error) {
	var test Test

	err := s.db.QueryRowContext(ctx, "select id, text from test where id = $1", id).Scan(
		&test.ID,
		&test.Text,
	)
	if err != nil {
		return Test{}, err
	}

	return test, nil
}
