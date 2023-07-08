package database

import (
	"coffeeshop/internal/orders/domain"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

/// REPOSITORY LAYER (sqlite adapter)

type Repo struct {
	*sql.DB
}

// Connect to DB
func Connect(dbPath string) (Repo, error) {
	db, err := sql.Open("sqlite3", dbPath)
	return Repo{db}, err
}

func (repo *Repo) GetCoffeeIdAndPriceByName(coffeeName string) (coffeeID uint32, price float32, err error) {
	err = repo.QueryRow(
		`SELECT coffeeID, price FROM coffee WHERE name = $1;`,
		coffeeName).Scan(&coffeeID, &price)
	return
}

func (repo *Repo) GetToppingIdAndPriceByName(toppingName string) (toppingID uint32, price float32, err error) {
	err = repo.QueryRow(
		`SELECT toppingID, price FROM topping WHERE name = $1;`,
		toppingName).Scan(&toppingID, &price)
	return
}

func (repo *Repo) CreateOrder(userID, coffeeID, toppingID, sugar uint32, date time.Time) (uint32, error) {
	res, err := repo.Exec(
		`INSERT INTO order_ (userID, coffeeID, toppingID, sugar, date) VALUES ($1, $2, $3, $4, $5);`,
		userID, coffeeID, toppingID, sugar, date)
	id, _ := res.LastInsertId()
	return uint32(id), err
}

func (repo *Repo) GetOrderById(userID, orderID uint32) (order *domain.Order, err error) {
	var toppingPrice float32
	err = repo.QueryRow(`
		SELECT coffee.name, coffee.price, topping.name, topping.price, sugar, status, date
		FROM order_ INNER JOIN
			coffee ON order_.coffeeID = coffee.coffeeID INNER JOIN
			topping ON order_.toppingID = topping.toppingID
		WHERE order_.userID = $1 AND order_.orderID = $2;
		`, userID, orderID).Scan(&order.Coffee, &order.Total, &order.Topping,
								 &toppingPrice, &order.Sugar, &order.Status, &order.Date)
	order.Total += toppingPrice
	return
}

func (repo *Repo) GetOrdersByShift(userID, shift, limit uint32) (orders []*domain.Order, err error) {
	rows, err := repo.Query(`
	SELECT orderID, coffee.name, coffee.price, topping.name, topping.price, sugar, status, date
	FROM order_ INNER JOIN
		coffee ON order_.coffeeID = coffee.coffeeID INNER JOIN
		topping ON order_.toppingID = topping.toppingID
	WHERE order_.userID = $1
	LIMIT $2 OFFSET $3;
	`, userID, limit, shift)
	if err != nil { return nil, err }
	defer rows.Close()
	// Append rows to orders slice
	var toppingPrice float32
	for rows.Next() {
		order := domain.Order{Userid: userID}
		if err := rows.Scan(&order.Id, &order.Coffee, &order.Total, &order.Topping,
							&toppingPrice, &order.Sugar, &order.Status, &order.Date);
		err != nil { return nil, err }
		order.Total += toppingPrice
		orders = append(orders, &order)
	}
	return
}

func (repo *Repo) SetStatusCancelled(userID, orderID uint32) (err error) {
	_, err = repo.Exec(
		`UPDATE order_ SET status = 2 WHERE userID = $1 AND orderID = $2;`,
		userID, orderID)
	return
}

func (repo *Repo) DeleteById(userID, orderID uint32) (err error) {
	_, err = repo.Exec(
		`DELETE FROM order_ WHERE userID = $1 AND orderID = $2;`,
		userID, orderID)
	return
}
