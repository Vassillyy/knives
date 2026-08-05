package dto

import (
	"time"

	"knives/internal/domain"
)

type KnifePhotoResponse struct {
	ID        string    `json:"id"`
	KnifeID   string    `json:"knife_id"`
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at"`
}

func KnifePhotoFromDomain(p *domain.KnifePhoto) KnifePhotoResponse {
	return KnifePhotoResponse{
		ID:        p.ID,
		KnifeID:   p.KnifeID,
		Filename:  p.Filename,
		CreatedAt: p.CreatedAt,
	}
}

func KnifePhotosFromDomain(photos []domain.KnifePhoto) []KnifePhotoResponse {
	out := make([]KnifePhotoResponse, len(photos))
	for i, p := range photos {
		out[i] = KnifePhotoFromDomain(&p)
	}
	return out
}
