package repositories

import (
	"AiCheto/internal/entity"
	"context"
	"github.com/google/uuid"
)

type IQuoteRepository interface {
	Create(ctx context.Context, q *entity.Quote) error
	GetAll(ctx context.Context) ([]entity.Quote, error)
	GetByAuthor(ctx context.Context, author string) ([]entity.Quote, error)
	GetRandom(ctx context.Context) (*entity.Quote, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
