package handler

import (
	"context"
	storageDto "example/internal/storage"
)

type storage interface {
	SelectTestByID(ctx context.Context, id int) (storageDto.Test, error)
	CreateTest(ctx context.Context, text string) error
}
