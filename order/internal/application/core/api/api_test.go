package api

import (
	"context"
	"testing"

	"github.com/asadlive84/microservices/order/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPayment struct {
	mock.Mock
}

func (p *mockPayment) Charge(ctx context.Context, payment *domain.Order) error {
	args := p.Called(payment)
	return args.Error(0)
}

type mockDb struct {
	mock.Mock
}

func (d *mockDb) Get(id int64) (domain.Order, error) {
	args := d.Called(id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func (d *mockDb) Save(order *domain.Order) error {
	args := d.Called(order)
	return args.Error(0)
}

func Test_Should_Palce_Order(t *testing.T) {
	payment := new(mockPayment)
	db := new(mockDb)

	payment.On("Charge", mock.Anything).Return(nil)
	db.On("Save", mock.Anything).Return(nil)

	app := NewApplication(db, payment)

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
}
