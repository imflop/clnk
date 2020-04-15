package sqlstore

import (
	"github.com/imflop/clnk/internal/app/models"
	uuid "github.com/satori/go.uuid"
	"time"
)

// LinkRepository ...
type LinkRepository struct {
	store *Store
}

// Create ...
func (r *LinkRepository) Create(originalURL string) (*models.Link, error) {
	l := &models.Link{}
	u := uuid.NewV4()
	datetime := time.Now().Format(time.RFC3339)
	err := r.store.db.QueryRow(
		"INSERT INTO core_link (uuid, original_url, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		u,
		originalURL,
		datetime,
		datetime,
	).Scan(&l.ID)
	if err != nil {
		return nil, err
	} else {
		return l, nil
	}
}

// Find ...
func (r *LinkRepository) Find(id int) (*models.Link, error) {
	l := &models.Link{}
	err := r.store.db.QueryRow(
		"SELECT uuid, original_url, short_url FROM core_link WHERE id=$1",
		id,
	).Scan(&l.UUID, &l.OriginalURL, &l.ShortURL)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Update ...
func (r *LinkRepository) Update(id int, shortURL string) (*models.Link, error) {
	l := &models.Link{}
	err := r.store.db.QueryRow(
		"UPDATE core_link SET short_url=$1 WHERE id=$2 RETURNING short_url",
		shortURL,
		id,
	).Scan(&l.ShortURL)
	if err != nil {
		return nil, err
	} else {
		return l, nil
	}
}
