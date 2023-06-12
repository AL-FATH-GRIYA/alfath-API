package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Brand struct {
	bun.BaseModel
	ID        int64     `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,unique"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
}

func (m *Brand) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
