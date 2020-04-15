package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Link ...
type Link struct {
	ID          int        `json:"id,omitempty"`
	UUID        *uuid.UUID `json:"uuid,omitempty"`
	OriginalURL string     `json:"original_url,omitempty"`
	ShortURL    string     `json:"short_url,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
