package ports

import (
	"context"
	"github.com/asadlive84/microservices/payment/internal/application/core/domain"
)

type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}