package db

import (
	"fmt"

	"github.com/asadlive84/microservices/order/internal/application/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, oepnErr := gorm.Open(postgres.Open(dataSourceUrl), &gorm.Config{})
	if oepnErr != nil {
		return nil, fmt.Errorf("db connection error: %v", oepnErr)
	}

	err := db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(id int64) (domain.Order, error) {
	var orderEntity Order
	res := a.db.First(&orderEntity, id)
	var oderItems []domain.OrderItem
	for _, orderItem := range orderEntity.OrderItems {
		oderItems = append(oderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: oderItems,
		CreatedAt:  orderEntity.CreatedAt.UnixNano(),
	}
	return order, res.Error
}

func (a *Adapter) Save(order *domain.Order) error {
	var orderItems []OrderItem
	for _, orderItem := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}

	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.Create(&orderModel)
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}
