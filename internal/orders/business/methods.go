package business

import (
	db "coffeeshop/internal/orders/database"
	"coffeeshop/internal/orders/domain"
	"errors"
	"log"
	"time"
)

/// APPLICATION BUSINESS LOGIC LAYER

type Business struct {
	repo *db.Repo
}

func New(repo *db.Repo) Business {
	return Business{repo: repo}
}

// Create order
func (b *Business) Create(userID, sugar uint32, coffee, topping string) (*domain.Order, error) {
	// Validation
	if coffee != "" {
		return nil, errors.New("you didn't specify any coffee")
	}
	// Sum coffee price & get id
	coffeeID, total, err := b.repo.GetCoffeeIdAndPriceByName(coffee)
	if err != nil {
		return nil, errors.New("specified coffee type was wrong")
	}
	// Sum topping price & get id
	var toppingID uint32
	if topping != "" {
		var toppingPrice float32
		toppingID, toppingPrice, err = b.repo.GetToppingIdAndPriceByName(topping)
		if err != nil {
			return nil, errors.New("specified topping type was wrong")
		}
		total += toppingPrice
	}
	// Append to DB
	date := time.Now()
	orderID, err := b.repo.CreateOrder(userID, coffeeID, toppingID, sugar, date)
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
func (b *Business) Get(userID, orderID uint32) (*domain.Order, error) {
	order, err := b.repo.GetOrder(orderID, userID)
	if err != nil {
		return nil, errors.New("specified order id was wrong")
	}
	order.Id = orderID
	order.Userid = userID
	return order, nil
}

// Get some orders by shift
func (b *Business) GetSome(userID, shift uint32) ([]*domain.Order, error) {
	orders, err := b.repo.GetSome(userID, shift)
	if err != nil {
		return nil, errors.New("specified order shift was wrong")
	}
	return orders, nil
}

// Cancel pending order
// func (s *OrdersServiceServer) Cancel(ctx context.Context, in *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
// 	// Get changing order
// 	getResponce, err := s.Get(ctx, &pb.GetOneOrderRequest{
// 		Id: in.GetId(),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Validate order status
// 	switch getResponce.Order.Status {
// 	case 1:
// 		return nil, errors.New("order was already DELIVERED")
// 	case 2:
// 		return nil, errors.New("order was already CANCELLED")
// 	}
// 	// Change order status to CANCELLED
// 	if _, err := s.db.Exec(
// 		`UPDATE order_ SET status = 2 WHERE orderID = $1 AND userID = $2;`,
// 		getResponce.Order.Id, getResponce.Order.Userid); err != nil {
// 		return nil, err
// 	}
// 	getResponce.Order.Status = 2
// 	return &pb.CancelOrderResponse{Order: getResponce.Order}, nil
// }

// // Delete order
// func (s *OrdersServiceServer) Delete(ctx context.Context, in *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
// 	// Validate input
// 	if in.GetId() == 0 {
// 		return nil, errors.New("order ID is wrong")
// 	}
// 	// Get deleting order
// 	getResponce, err := s.Get(ctx, &pb.GetOneOrderRequest{
// 		Id: in.GetId(),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Delete
// 	if _, err := s.db.Exec(
// 		`DELETE FROM order_ WHERE orderID = $1 AND userID = $2;`,
// 		getResponce.Order.Id, getResponce.Order.Userid); err != nil {
// 		return nil, err
// 	}
// 	return &pb.DeleteOrderResponse{Order: getResponce.Order}, nil
// }
