package post

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (pg *Postgres) CreateDelivery(ctx context.Context, delivery Delivery) (int64, error) {
	query := `
	INSERT INTO delivery (name, phone, zip, city, address, region, email) 
	VALUES (@name, @phone, @zip, @city, @address, @region, @email) RETURNING id`

	args := pgx.NamedArgs{
		"name":    delivery.Name,
		"phone":   delivery.Phone,
		"zip":     delivery.Zip,
		"city":    delivery.City,
		"address": delivery.Address,
		"region":  delivery.Region,
		"email":   delivery.Email,
	}

	result := pg.db.QueryRow(ctx, query, args)

	var id int64
	err := result.Scan(&id)

	if err != nil {
		return -1, fmt.Errorf("unable to insert row: %w", err)
	}

	return id, nil
}

func (pg *Postgres) GetDelivery(ctx context.Context, delivery_id int64) (*Delivery, error) {
	query := `
	select * 
	from delivery
	where id = @id`

	args := pgx.NamedArgs{
		"id": delivery_id,
	}

	rows, err := pg.db.Query(ctx, query, args)

	if err != nil {
		return &Delivery{}, fmt.Errorf("unable to query delivery: %w", err)
	}

	defer rows.Close()
	d, err := pgx.CollectRows(rows, pgx.RowToStructByName[Delivery])

	if len(d) < 1 {
		return &Delivery{}, fmt.Errorf("not delivery on id: %d", delivery_id)
	}

	if err != nil {
		return &Delivery{}, fmt.Errorf("unable to query delivery: %w", err)
	}

	return &d[0], err
}
