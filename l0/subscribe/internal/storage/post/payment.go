package post

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Payment struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        int64  `json:"amount"`
	Payment_dt    int64  `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost int64  `json:"delivery_cost"`
	Goods_total   int64  `json:"goods_total"`
	Custom_fee    int64  `json:"custom_fee"`
}

func (pg *Postgres) CreatePayment(ctx context.Context, payment Payment) (int64, error) {
	query := `
	INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
	VALUES (@transaction, @request_id, @currency, @provider, @amount, @payment_dt, @bank, @delivery_cost, @goods_total, @custom_fee) RETURNING id`

	args := pgx.NamedArgs{
		"transaction":   payment.Transaction,
		"request_id":    payment.Request_id,
		"currency":      payment.Currency,
		"provider":      payment.Provider,
		"amount":        payment.Amount,
		"payment_dt":    payment.Payment_dt,
		"bank":          payment.Bank,
		"delivery_cost": payment.Delivery_cost,
		"goods_total":   payment.Goods_total,
		"custom_fee":    payment.Custom_fee,
	}

	result := pg.db.QueryRow(ctx, query, args)

	var id int64
	err := result.Scan(&id)

	if err != nil {
		return -1, fmt.Errorf("unable to insert row: %w", err)
	}

	return id, nil
}

func (pg *Postgres) GetPayment(ctx context.Context, payment_id int64) (*Payment, error) {
	query := `
	select * 
	from payment
	where id = @id`

	args := pgx.NamedArgs{
		"id": payment_id,
	}

	rows, err := pg.db.Query(ctx, query, args)

	if err != nil {
		return &Payment{}, fmt.Errorf("unable to query  payment: %w", err)
	}

	defer rows.Close()
	d, err := pgx.CollectRows(rows, pgx.RowToStructByName[Payment])

	if len(d) < 1 {
		return &Payment{}, fmt.Errorf("not  payment on id: %d", payment_id)
	}

	if err != nil {
		return &Payment{}, fmt.Errorf("unable to query  payment: %w", err)
	}

	return &d[0], err
}
