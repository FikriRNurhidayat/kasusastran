package repository

import (
	"context"

	"github.com/fikrirnurhidayat/kasusastran/app/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(context.Context, *entity.User) (*entity.User, error)
	Get(context.Context, uuid.UUID) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	EmailExist(context.Context, string) (bool, error)
}
