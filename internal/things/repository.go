package things

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
)

type Repository interface {
	Create(ctx context.Context, thing Thing) error
	Load(ctx context.Context, id ID) (*Thing, error)
	LoadAll(ctx context.Context) ([]Thing, error)
}

func NewRepository(logger zerolog.Logger) Repository {
	return inMemoryRepo{
		logger: logger,
		data:   map[ID]Thing{},
	}
}

type inMemoryRepo struct {
	logger zerolog.Logger
	data   map[ID]Thing
}

var ErrAlreadyExists = errors.New("already exists")
var ErrNotFound = errors.New("not found")

func (repo inMemoryRepo) Create(ctx context.Context, thing Thing) error {
	if _, exists := repo.data[thing.ID]; exists {
		return ErrAlreadyExists
	}

	repo.data[thing.ID] = thing

	return nil
}

func (repo inMemoryRepo) Load(ctx context.Context, id ID) (*Thing, error) {
	thing, exists := repo.data[id]

	if !exists {
		return nil, ErrNotFound
	}

	return &thing, nil
}

func (repo inMemoryRepo) LoadAll(ctx context.Context) ([]Thing, error) {
	things := []Thing{}
	for _, thing := range repo.data {
		things = append(things, thing)
	}
	return things, nil
}
