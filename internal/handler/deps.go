package handler

import (
	"context"

	storageDto "example/internal/storage"
)

type storage interface {
	GetTestByID(ctx context.Context, id int) (storageDto.Test, error)
	GetTests(ctx context.Context) ([]storageDto.Test, error)
	CreateTest(ctx context.Context, text string) error
	GetRole(ctx context.Context, token string) (string, error)
}
