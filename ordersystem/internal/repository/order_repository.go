package repository

import (
	"database/sql"

	"github.com/cgalimberti/2099-CleanArch/internal/domain"
)

type OrderRepository interface {
	Create(order domain.Order) error
	List() ([]domain.Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order domain.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (id, customer_id, total, status) VALUES (?, ?, ?, ?)", order.ID, order.CustomerID, order.Total, order.Status)
	return err
}

func (r *orderRepository) List() ([]domain.Order, error) {
	rows, err := r.db.Query("SELECT id, customer_id, total, status FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.Total, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
