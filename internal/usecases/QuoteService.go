package usecases

import (
	"AiCheto/internal/entity"
	"AiCheto/internal/repositories"
	"context"
	"github.com/google/uuid"
)

type QuoteService struct {
	repo repositories.IQuoteRepository
}

func NewQuoteUsecase(r repositories.IQuoteRepository) *QuoteService {
	return &QuoteService{repo: r}
}

func (u *QuoteService) Create(ctx context.Context, q *entity.Quote) error {
	return u.repo.Create(ctx, q)
}

func (u *QuoteService) List(ctx context.Context, author string) ([]entity.Quote, error) {
	if author == "" {
		return u.repo.GetAll(ctx)
	}
	return u.repo.GetByAuthor(ctx, author)
}

func (u *QuoteService) Random(ctx context.Context) (*entity.Quote, error) {
	return u.repo.GetRandom(ctx)
}

func (u *QuoteService) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
