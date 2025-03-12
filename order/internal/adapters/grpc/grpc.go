package grpc

import (
	"context"
	"fmt"

	"github.com/asadlive84/microservices-proto-asad/golang/order"
	"github.com/asadlive84/microservices/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {

	fmt.Println("==============Start Create==========================")
	var orderItems []domain.OrderItem
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.Name,
			// UnitPrice:   orderItem.UnitPrice,
			// Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(ctx, &newOrder)
	if err != nil {
		fmt.Println("==============##2 Create==========================", err)
		return nil, err
	}
	fmt.Println("==============##3 result==========================", result)
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}
