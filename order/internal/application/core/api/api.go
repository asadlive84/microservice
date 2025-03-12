package api

import (
	"context"

	"github.com/asadlive84/microservices/order/internal/application/core/domain"
	"github.com/asadlive84/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
	}
}

func (a *Application) PlaceOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	err := a.db.Save(order)
	if err != nil {
		return &domain.Order{}, err
	}

	// ctx1, _ := context.WithTimeout(context.Background(), 16*time.Second)

	// fmt.Println("============= time out pattern =======")
	// fmt.Printf("%+v\n", ctx1)
	// fmt.Println("======================================")

	paymentErr := a.payment.Charge(ctx, order)
	if paymentErr != nil {
		st, _ := status.FromError(paymentErr)
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: st.Message(),
		}

		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = append(badRequest.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation faild on payment failed")
		statusDetails, _ := orderStatus.WithDetails(badRequest)

		return &domain.Order{}, statusDetails.Err()
	}

	return order, nil
}
