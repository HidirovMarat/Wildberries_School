package post

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Order struct {
	Order_uid          string    `json:"order_uid"`
	Track_number       string    `json:"track_number"`
	Entry              string    `json:"entry"`
	Delivery           Delivery  `json:"delivery"`
	Payment            Payment   `json:"payment"`
	Items              []Item    `json:"items"`
	Locale             string    `json:"locale"`
	Internal_signature string    `json:"internal_signature"`
	Customer_id        string    `json:"customer_id"`
	Delivery_service   string    `json:"delivery_service"`
	Shardkey           string    `json:"shardkey"`
	Sm_id              int64     `json:"sm_id"`
	Date_created       time.Time `json:"date_created"`
	Oof_shard          string    `json:"oof_shard"`
}

type OrderDB struct {
	Order_uid          string    `json:"order_uid"`
	Track_number       string    `json:"track_number"`
	Entry              string    `json:"entry"`
	Delivery_id        int64     `json:"delivery_id"`
	Payment_id         int64     `json:"payment_id"`
	Item_id            int64     `json:"item_id"`
	Locale             string    `json:"locale"`
	Internal_signature string    `json:"internal_signature"`
	Customer_id        string    `json:"customer_id"`
	Delivery_service   string    `json:"delivery_service"`
	Shardkey           string    `json:"shardkey"`
	Sm_id              int64     `json:"sm_id"`
	Date_created       time.Time `json:"date_created"`
	Oof_shard          string    `json:"oof_shard"`
}

func (pg *Postgres) CreateOrder(ctx context.Context, order Order) error {
	delivery_id, err := pg.CreateDelivery(ctx, order.Delivery)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	payment_id, err := pg.CreatePayment(ctx, order.Payment)

	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	for _, item := range order.Items {
		item_id, err := pg.CreateItem(ctx, item)

		if err != nil {
			log.Printf("unable to insert row:")
			continue
		}

		query := `
				INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, item_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
				VALUES (@order_uid, @track_number, @entry, @delivery_id, @payment_id, @item_id, @locale, @internal_signature, @customer_id, @delivery_service, @shardkey, @sm_id, @date_created, @oof_shard)
			`

		args := pgx.NamedArgs{
			"order_uid":          order.Order_uid,
			"track_number":       order.Track_number,
			"entry":              order.Entry,
			"delivery_id":        delivery_id,
			"payment_id":         payment_id,
			"item_id":            item_id,
			"locale":             order.Locale,
			"internal_signature": order.Internal_signature,
			"customer_id":        order.Customer_id,
			"delivery_service":   order.Delivery_service,
			"shardkey":           order.Shardkey,
			"sm_id":              order.Sm_id,
			"date_created":       order.Date_created,
			"oof_shard":          order.Oof_shard,
		}

		_, err = pg.db.Exec(ctx, query, args)

		if err != nil {
			return fmt.Errorf("unable to insert row: %w", err)
		}
	}

	return nil
}

func (pg *Postgres) GetOrder(ctx context.Context, order_id string) (*Order, error) {
	query := `
	select * 
	from orders
	where order_uid = @id`

	args := pgx.NamedArgs{
		"id": order_id,
	}

	rows, err := pg.db.Query(ctx, query, args)

	if err != nil {
		return &Order{}, fmt.Errorf("unable to query order: %w", err)
	}

	defer rows.Close()
	d, err := pgx.CollectRows(rows, pgx.RowToStructByName[OrderDB])

	if len(d) < 1 {
		return &Order{}, fmt.Errorf("not order on id: %s", order_id)
	}

	if err != nil {
		return &Order{}, fmt.Errorf("unable to query order: %w", err)
	}

	orderDB := d[0]
	order := Order{}

	delivery, err := pg.GetDelivery(ctx, orderDB.Delivery_id)

	if err != nil {
		return &Order{}, fmt.Errorf("not  delivery on id: %s", order_id)
	}

	payment, err := pg.GetPayment(ctx, orderDB.Payment_id)

	if err != nil {
		return &Order{}, fmt.Errorf("not  payment on id: %s", order_id)
	}

	order.Delivery = *delivery
	order.Payment = *payment

	for _, orderVal := range d {
		item, err := pg.GetItem(ctx, orderVal.Item_id)

		if err != nil {
			return &Order{}, fmt.Errorf("not  item on id: %s", order_id)
		}

		order.Items = append(order.Items, *item)
	}

	return &order, err
}

func (pg *Postgres) GetAll(ctx context.Context) (*[]Order, error) {
	uids, err := pg.getUIDs(ctx)

	if err != nil {
		return &[]Order{}, err
	}

	orders := []Order{}

	for _, uid := range *uids {
		order, err := pg.GetOrder(ctx, uid)

		if err != nil {
			return &[]Order{}, err
		}

		orders = append(orders, *order)
	}

	return &orders, nil
}

func (pg *Postgres) getUIDs(ctx context.Context) (*[]string, error) {
	query := `
	select order_uid 
	from orders`

	rows, err := pg.db.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("unable to query order: %w", err)
	}

	defer rows.Close()

	var uids []string
	for rows.Next() {
		var uid string
		if err := rows.Scan(&uid); err != nil {
			return nil, fmt.Errorf("unable to scan order_uid: %w", err)
		}
		uids = append(uids, uid)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return &uids, nil
}
