package get

import (
	"net/http"

	"log/slog"

	"context"
	"sub/internal/storage/post"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type OrderGetter interface {
	GetOrder(ctx context.Context, order_id string) (*post.Order, error)
}

type Response struct {
	Order post.Order `json:"order,omitempty"`
}

func New(ctx context.Context, log *slog.Logger, cash map[string]post.Order, orderGetter OrderGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.order.get.New"

		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		order_uid := r.URL.Query().Get("id")

		if order_uid == "" {
			log.Info("id empty in url")
			http.Error(w, "id empty", http.StatusBadRequest)
			return
		}

		if orderCash, ok := cash[order_uid]; ok {
			log.Info("get order at cash", "id", orderCash.Order_uid)
			responseOK(w, r, orderCash)
			return
		}

		orderDB, err := orderGetter.GetOrder(ctx, order_uid)

		if err != nil {
			log.Info("Not id", "id", orderDB.Order_uid)
			http.Error(w, "id not", http.StatusBadRequest)
			return
		}

		responseOK(w, r, *orderDB)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, order post.Order) {
	render.JSON(w, r, Response{
		Order: order,
	})
}
