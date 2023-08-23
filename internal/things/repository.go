package things

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/rs/zerolog"
)

type Repository interface {
	Create(thing Thing) error
	Load(id ID) (*Thing, error)
	LoadAll() ([]Thing, error)
}

var ErrAlreadyExists = errors.New("already exists")
var ErrNotFound = errors.New("not found")

func NewRepository(logger zerolog.Logger, db *goqu.Database) Repository {
	return dbRepo{
		logger: logger,
		db:     db,
	}
}

type dbRepo struct {
	logger zerolog.Logger
	db     *goqu.Database
}

func (repo dbRepo) Create(thing Thing) error {
	_, err := repo.db.Insert("thing").Rows(thing).Executor().Exec()
	return err
}

func (repo dbRepo) Load(id ID) (*Thing, error) {
	rows, err := repo.filter(filter{id: &id})

	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, ErrNotFound
	}

	return &rows[0], nil
}

func (repo dbRepo) LoadAll() ([]Thing, error) {
	return repo.filter(filter{})
}

// filter could contain more properties
type filter struct {
	id *ID
}

// This is private in this example but could be made public for more flexibility
func (repo dbRepo) filter(filter filter) ([]Thing, error) {
	where := goqu.Ex{}

	if filter.id != nil {
		where["id"] = filter.id
	}

	ds := repo.db.
		From("thing").
		Where(where).
		Order(goqu.C("name").Asc())

	rows := []Thing{}
	err := ds.ScanStructs(&rows)

	return rows, err
}

func NewInMemoryRepository(logger zerolog.Logger) Repository {
	return inMemoryRepo{
		logger: logger,
		data:   map[ID]Thing{},
	}
}

type inMemoryRepo struct {
	logger zerolog.Logger
	data   map[ID]Thing
}

func (repo inMemoryRepo) Create(thing Thing) error {
	if _, exists := repo.data[thing.ID]; exists {
		return ErrAlreadyExists
	}

	repo.data[thing.ID] = thing

	return nil
}

func (repo inMemoryRepo) Load(id ID) (*Thing, error) {
	thing, exists := repo.data[id]

	if !exists {
		return nil, ErrNotFound
	}

	return &thing, nil
}

func (repo inMemoryRepo) LoadAll() ([]Thing, error) {
	things := []Thing{}
	for _, thing := range repo.data {
		things = append(things, thing)
	}
	return things, nil
}
