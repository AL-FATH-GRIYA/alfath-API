package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Branch struct {
	bun.BaseModel
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	Name      string    `bun:"name,unique" json:"name"`
	Address   string    `bun:"address" json:"address"`
	CreatedAt time.Time `bun:"created_at" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at" json:"updated_at"`
}

func (m *Branch) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
