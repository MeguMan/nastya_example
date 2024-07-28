package storage

import (
	"context"
)

func (s *Storage) GetRole(ctx context.Context, token string) (string, error) {
	var role string
	err := s.db.QueryRowContext(ctx, "select role from person where token = $1", token).Scan(
		&role,
	)
	if err != nil {
		return role, err
	}

	return role, nil
}
