package cash

import (
	"context"
	"fmt"
	"log"
	"sub/internal/storage/post"
)

func New(ctx context.Context, pg *post.Postgres) (map[string]post.Order, error) {
	orders, err := pg.GetAll(ctx)
	cash := make(map[string]post.Order)

	if err != nil {
		return cash, fmt.Errorf("error orders of cash")
	}

	if orders == nil {
		return cash, nil
	}

	for _, order := range *orders {
		cash[order.Order_uid] = order
	}
	log.Print("good cash")
	return cash, nil
}
