package store

import "github.com/imflop/clnk/internal/app/models"

// LinkRepository ...
type LinkRepository interface {
	Find(int) (*models.Link, error)
	Create(string) (*models.Link, error)
	Update(int, string) (*models.Link, error)
}
