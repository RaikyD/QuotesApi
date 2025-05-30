package repositories

import (
	"context"
	"errors"
	"github.com/RaikyD/QuotesApi/internal/entity"
	"github.com/google/uuid"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func init() { rand.Seed(time.Now().UnixNano()) }

type InMemoryQuoteRepository struct {
	mu     sync.RWMutex
	quotes map[uuid.UUID]entity.Quote
}

func New() *InMemoryQuoteRepository {
	return &InMemoryQuoteRepository{quotes: make(map[uuid.UUID]entity.Quote)}
}

func (r *InMemoryQuoteRepository) Create(_ context.Context, q *entity.Quote) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	if _, exists := r.quotes[q.ID]; exists {
		return errors.New("duplicate id")
	}

	r.quotes[q.ID] = *q
	return nil
}

func (r *InMemoryQuoteRepository) GetAll(_ context.Context) ([]entity.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]entity.Quote, 0, len(r.quotes))
	for _, q := range r.quotes {
		out = append(out, q)
	}
	return out, nil
}

func (r *InMemoryQuoteRepository) GetByAuthor(_ context.Context, author string) ([]entity.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var out []entity.Quote
	for _, q := range r.quotes {
		if strings.EqualFold(q.Author, author) { // регистр уже не важен
			out = append(out, q)
		}
	}
	return out, nil
}

func (r *InMemoryQuoteRepository) GetRandom(_ context.Context) (*entity.Quote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.quotes) == 0 {
		return nil, errors.New("no quotes available")
	}
	idx := rand.Intn(len(r.quotes))
	i := 0
	for _, q := range r.quotes {
		if i == idx {
			return &q, nil
		}
		i++
	}
	return nil, errors.New("random selection failed")
}

func (r *InMemoryQuoteRepository) Delete(_ context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.quotes[id]; !ok {
		return errors.New("quote not found")
	}
	delete(r.quotes, id)
	return nil
}
