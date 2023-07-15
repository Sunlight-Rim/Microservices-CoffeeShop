package business

import (
	"coffeeshop/internal/orders/domain"
	"errors"
	"log"
	"time"
)

/// APPLICATION BUSINESS LOGIC LAYER

type Repository interface {
	GetCoffeeIdAndPriceByName(string) (uint32, float32, error)
	GetToppingIdAndPriceByName(string) (uint32, float32, error)
	CreateOrder(uint32, uint32, uint32, uint32, time.Time) (uint32, error)
	GetOrderById(uint32, uint32) (*domain.Order, error)
	GetOrdersByShift(uint32, uint32, uint32) ([]*domain.Order, error)
	SetStatusCancelled(uint32, uint32) (error)
	DeleteById(uint32, uint32) (error)
}

type Logic struct {
	repo Repository
}

func New(repo Repository) *Logic {
	return &Logic{repo: repo}
}

// Create order
func (l *Logic) Create(userID, sugar uint32, coffee, topping string) (*domain.Order, error) {
	// Validation
	if coffee == "" {
		return nil, errors.New("you didn't specify any coffee")
	}
	// Sum coffee price & get id
	coffeeID, total, err := l.repo.GetCoffeeIdAndPriceByName(coffee)
	if err != nil {
		return nil, errors.New("specified coffee type was wrong")
	}
	// Sum topping price & get id
	var toppingID uint32
	if topping != "" {
		var toppingPrice float32
		toppingID, toppingPrice, err = l.repo.GetToppingIdAndPriceByName(topping)
		if err != nil {
			return nil, errors.New("specified topping type was wrong")
		}
		total += toppingPrice
	}
	// Append to DB
	date := time.Now()
	orderID, err := l.repo.CreateOrder(userID, coffeeID, toppingID, sugar, date)
	if err != nil {
		log.Printf("DB request error: %v", err)
		return nil, errors.New("there is some problem with DB")
	}
	return &domain.Order{
		Id:      orderID,
		Userid:  userID,
		Coffee:  coffee,
		Topping: topping,
		Sugar:   sugar,
		Date:    date,
		Status:  0,
		Total:   total,
	}, nil
}

// Get one order by id
func (l *Logic) Get(userID, orderID uint32) (*domain.Order, error) {
	order, err := l.repo.GetOrderById(orderID, userID)
	if err != nil {
		return nil, errors.New("specified order id was wrong")
	}
	order.Id = orderID
	order.Userid = userID
	return order, nil
}

// Get some orders by shift
func (l *Logic) GetSome(userID, shift uint32) ([]*domain.Order, error) {
	var limit uint32 = 5
	orders, err := l.repo.GetOrdersByShift(userID, shift, limit)
	if err != nil {
		return nil, errors.New("specified order shift was wrong")
	}
	return orders, nil
}

// Cancel pending order
func (l *Logic) Cancel(userID, orderID uint32) (*domain.Order, error) {
	// Get changing order
	order, err := l.Get(userID, orderID)
	if err != nil {
		return nil, err
	}
	// Validate order status is PENDING
	switch order.Status {
	case 1:
		return nil, errors.New("order was already DELIVERED")
	case 2:
		return nil, errors.New("order was already CANCELLED")
	}
	// Change order status to CANCELLED
	if err = l.repo.SetStatusCancelled(userID, orderID); err != nil {
		return nil, err
	}
	order.Status = 2
	return order, nil
}

// Delete order
func (l *Logic) Delete(userID, orderID uint32) (*domain.Order, error) {
	// Get deleting order
	order, err := l.Get(userID, orderID)
	if err != nil {
		return nil, err
	}
	// Delete
	l.repo.DeleteById(userID, orderID)
	return order, nil
}
