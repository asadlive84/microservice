package api

import (
	"context"
	"testing"

	"github.com/asadlive84/microservices/order/internal/application/core/domain"
	// "github.com/asadlive84/microservices/order/internal/ports"
	"github.com/asadlive84/microservices/order/mocks/internal_/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPlaceOrder_Success(t *testing.T) {
    // ১. প্রস্তুতি (Arrange)
    mockPayment := &mocks.PaymentPort{}
    mockDB := &mocks.DBPort{}

    // mockPayment.On("Charge", mock.Anything).Return(nil)
    // mockDB.On("Save", mock.Anything).Return(nil)

	ctx:=context.Background()

	mockDB.On("Save", mock.AnythingOfType("*domain.Order")).Return(nil)
    mockPayment.On("Charge", ctx, mock.AnythingOfType("*domain.Order")).Return(nil)

    app := NewApplication(mockDB, mockPayment)

    _, err := app.PlaceOrder(context.Background(), &domain.Order{
        ID: 1,
        OrderItems: []domain.OrderItem{
            {
                ProductCode: "p1",
                // UnitPrice:   10,
                Quantity: 2,
            },
        },
        CreatedAt: 0,
    })
    assert.Nil(t, err)
    mockPayment.AssertExpectations(t)
    mockDB.AssertExpectations(t)
}
