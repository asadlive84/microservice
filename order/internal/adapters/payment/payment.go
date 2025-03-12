package payment

import (
	"context"

	pb "github.com/asadlive84/microservices-proto-asad/golang/payment"
	"github.com/asadlive84/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment pb.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	conn, err := grpc.NewClient(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := pb.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(ctx context.Context, order *domain.Order) error {
	_, err := a.payment.Create(ctx, &pb.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	if err != nil {
		return err
	}
	return nil
}
