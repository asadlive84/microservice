package ports

import (
	"context"

	"github.com/asadlive84/microservices/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(ctx context.Context,order *domain.Order) (*domain.Order, error)
}

type DBPort interface {
	Get(id int64) (domain.Order, error)
	Save(order *domain.Order) error
}
