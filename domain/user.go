package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type userRole string

const (
	SuperAdmin userRole = "super admin"
	admin      userRole = "admin"
)

type User struct {
	bun.BaseModel
	ID        int64     `bun:"id,pk,autoincrement"`
	Email     string    `bun:"email,unique"`
	Password  string    `bun:"password"`
	Role      userRole  `bun:"role"`
	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
}

func (m *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
