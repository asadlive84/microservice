package grpc

import (
	"context"
	"fmt"
	pb "github.com/asadlive84/microservices-proto-asad/golang/payment"
	"github.com/asadlive84/microservices/payment/internal/application/core/domain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	log.WithContext(ctx).Info("Creating payment...")
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &pb.CreatePaymentResponse{PaymentId: result.ID}, nil
}