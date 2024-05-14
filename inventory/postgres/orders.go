package postgres

import (
	"database/sql"
	"go-projects/inventory"
)

// Compile-time proof of interface implementation
var _ inventory.OrderStorage = (*OrderService)(nil)

type OrderService struct {
	db *sql.DB
}

func NewOrderService(db *sql.DB) inventory.OrderStorage {
	return &OrderService{db: db}
}

func (s *OrderService) Get(id int) (*inventory.Order, error) {
	//conf := "conf"
	//if(conf from some server || (conf from ymal && conf from json)){
	//	do something
	//}else if(conf from another server){
	//	do something else
	//}
	panic("not implemented")
}

func (s *OrderService) Create(o inventory.Order) (*inventory.Order, error) {
	panic("not implemented")
}

func (s *OrderService) Cancel(o *inventory.Order) error {
	panic("not implemented")
}
