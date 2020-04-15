package sqlstore

import (
	"github.com/imflop/clnk/internal/app/store"
	"github.com/jmoiron/sqlx"
)

// Store implementation ...
type Store struct {
	db             *sqlx.DB
	linkRepository *LinkRepository
}

// New ... constructor for store
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// Link ...
func (s *Store) Link() store.LinkRepository {
	if s.linkRepository != nil {
		return s.linkRepository
	}
	s.linkRepository = &LinkRepository{
		store: s,
	}
	return s.linkRepository
}
