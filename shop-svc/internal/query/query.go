package query

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type QueryInit struct {
	Db *sqlx.DB
}

var NotFound = errors.New("not found")

type Shop struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	Address   string       `db:"address"`
	OwnerID   string       `db:"owner_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type Query interface {
	CreateShop(ctx context.Context, shop Shop) (Shop, error)
	GetShopByID(ctx context.Context, shopID string) (Shop, error)
	UpdateShop(ctx context.Context, shop Shop) (Shop, error)
	DeleteShop(ctx context.Context, shopID string) error
	ListShops(ctx context.Context, ownerID string) ([]Shop, error)
	GetShopByName(ctx context.Context, shopName string) (Shop, error)
}
