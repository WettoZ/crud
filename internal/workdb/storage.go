package workdb

import (
	"context"
)

type Ways interface {
	InsertRow(ctx context.Context, us *UserData) error
	DeleteRow(ctx context.Context, element int64) error
	AllRows(ctx context.Context) ([]*UserData, error)
}
