package query

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// CreateShop creates a new shop in the database
func (q *QueryInit) CreateShop(ctx context.Context, s Shop) (Shop, error) {
	tx, err := q.Db.Beginx()
	if err != nil {
		return Shop{}, err
	}

	stmt, err := tx.PrepareNamed(`
        INSERT INTO shop (
            id, name, address, owner_id, created_at
        ) VALUES (
            uuid_generate_v4(), :name, :address, :owner_id, :created_at
        )
        RETURNING id, name, address, owner_id, created_at
    `)
	if err != nil {
		tx.Rollback()
		return Shop{}, err
	}

	s.CreatedAt = time.Now()
	// s.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	var newShop Shop
	err = stmt.Get(&newShop, &s)
	if err != nil {
		tx.Rollback()
		return Shop{}, err
	}

	err = tx.Commit()
	if err != nil {
		return Shop{}, err
	}

	return newShop, nil
}

func (q *QueryInit) GetShop(ctx context.Context, shopID string) (Shop, error) {
	var shop Shop
	err := q.Db.Get(&shop, `
        SELECT id, name, address, owner_id, created_at, updated_at
        FROM shop
        WHERE id = $1
    `, shopID)
	if err != nil {
		return Shop{}, err
	}
	return shop, nil
}

func (q *QueryInit) UpdateShop(ctx context.Context, s Shop) (Shop, error) {
	tx, err := q.Db.Beginx()
	if err != nil {
		return Shop{}, err
	}

	stmt, err := tx.PrepareNamed(`
        UPDATE shop
        SET name = :name, address = :address, owner_id = :owner_id, updated_at = :updated_at
        WHERE id = :id
        RETURNING id, name, address, owner_id, created_at, updated_at
    `)
	if err != nil {
		tx.Rollback()
		return Shop{}, err
	}

	s.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	var updatedShop Shop
	err = stmt.Get(&updatedShop, &s)
	if err != nil {
		tx.Rollback()
		return Shop{}, err
	}

	err = tx.Commit()
	if err != nil {
		return Shop{}, err
	}

	return updatedShop, nil
}

func (q *QueryInit) DeleteShop(ctx context.Context, shopID string) error {
	tx, err := q.Db.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
        DELETE FROM shop
        WHERE id = $1
    `, shopID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *QueryInit) GetShopByName(ctx context.Context, shopName string) (Shop, error) {
	var shop Shop

	query := `
		SELECT id, name, owner_id, address, created_at
		FROM shop
		WHERE name = $1
		LIMIT 1
	`

	err := q.Db.GetContext(ctx, &shop, query, shopName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Shop{}, nil // Return nil error to indicate not found
		}
		return Shop{}, err
	}

	return shop, nil
}

func (q *QueryInit) GetShopByID(ctx context.Context, shopID string) (Shop, error) {
	var shop Shop

	query := `
		SELECT id, name, owner_id, address, created_at
		FROM shop
		WHERE id = $1
		LIMIT 1
	`

	err := q.Db.GetContext(ctx, &shop, query, shopID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Shop{}, nil // Return nil error to indicate not found
		}
		return Shop{}, err
	}

	return shop, nil
}

func (q *QueryInit) ListShops(ctx context.Context, ownerID string) ([]Shop, error) {
	query := `
        SELECT id, name, owner_id, created_at
        FROM shop
        WHERE owner_id = $1
    `

	var shops []Shop
	err := q.Db.SelectContext(ctx, &shops, query, ownerID)
	if err != nil {
		return nil, err
	}

	return shops, nil
}
