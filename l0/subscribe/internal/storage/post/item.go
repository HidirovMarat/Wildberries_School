package post

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Item struct {
	Chrt_id      int64  `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int64  `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int64  `json:"sale"`
	Size         string `json:"size"`
	Total_price  int64  `json:"total_price"`
	Nm_id        int64  `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int64  `json:"status"`
}

func (pg *Postgres) CreateItem(ctx context.Context, item Item) (int64, error) {
	query := `
	INSERT INTO item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
	VALUES (@chrt_id, @track_number, @price, @rid, @name, @sale, @size, @total_price, @nm_id, @brand, @status) RETURNING id`

	args := pgx.NamedArgs{
		"chrt_id":      item.Chrt_id,
		"track_number": item.Track_number,
		"price":        item.Price,
		"rid":          item.Rid,
		"name":         item.Name,
		"sale":         item.Sale,
		"size":         item.Size,
		"total_price":  item.Total_price,
		"nm_id":        item.Nm_id,
		"brand":        item.Brand,
		"status":       item.Status,
	}

	result := pg.db.QueryRow(ctx, query, args)

	var id int64
	err := result.Scan(&id)

	if err != nil {
		return -1, fmt.Errorf("unable to insert row: %w", err)
	}

	return id, nil
}

func (pg *Postgres) GetItem(ctx context.Context, item_id int64) (*Item, error) {
	query := `
	select * 
	from item
	where id = @id`

	args := pgx.NamedArgs{
		"id": item_id,
	}

	rows, err := pg.db.Query(ctx, query, args)

	if err != nil {
		return &Item{}, fmt.Errorf("unable to query item: %w", err)
	}

	defer rows.Close()
	d, err := pgx.CollectRows(rows, pgx.RowToStructByName[Item])

	if len(d) < 1 {
		return &Item{}, fmt.Errorf("not item on id: %d", item_id)
	}

	if err != nil {
		return &Item{}, fmt.Errorf("unable to query item: %w", err)
	}

	return &d[0], err
}
