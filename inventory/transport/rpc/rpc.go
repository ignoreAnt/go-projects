package rpc

import (
	"go-projects/inventory"
	"go-projects/inventory/transport"
	"net"
	"net/rpc"
)

// Compile-time proof of interface implementation
var _ transport.InventoryTransporter = (*Service)(nil)

type Service struct {
	orderStore      inventory.OrderStorage
	productStore    inventory.ProductStorage
	supplierStore   inventory.SupplierStorage
	supplierService inventory.SupplierService
}

func NewRPCService(orderStore inventory.OrderStorage, supplierStore inventory.SupplierStorage, supplierService inventory.SupplierService, productStore inventory.ProductStorage) *Service {
	return &Service{
		orderStore:      orderStore,
		productStore:    productStore,
		supplierStore:   supplierStore,
		supplierService: supplierService,
	}

}

func (svc *Service) Serve(l net.Listener) error {
	err := rpc.Register(svc)
	if err != nil {
		return err
	}
	rpc.Accept(l) // blocks
	return nil
}

func (svc *Service) GetOrder(inventory.GetOrderRequest, *inventory.GetOrderResponse) error {
	panic("not implemented")
}

func (svc *Service) CreateOrder(inventory.CreateOrderRequest, *inventory.CreateOrderResponse) error {
	panic("not implemented")
}

func (svc *Service) OrderStatus(inventory.OrderStatusRequest, *inventory.OrderStatusResponse) error {
	panic("not implemented")
}

func (svc *Service) CancelOrder(inventory.CancelOrderRequest, *inventory.CancelOrderResponse) error {
	panic("not implemented")
}

func (svc *Service) GetProduct(inventory.GetProductRequest, *inventory.GetProductResponse) error {
	panic("not implemented")
}

func (svc *Service) CreateProduct(inventory.CreateProductRequest, *inventory.CreateProductResponse) error {
	panic("not implemented")
}

func (svc *Service) UpdateProduct(inventory.UpdateProductRequest, *inventory.UpdateProductResponse) error {
	panic("not implemented")
}

func (svc *Service) DeleteProduct(inventory.DeleteProductRequest, *inventory.DeleteProductResponse) error {
	panic("not implemented")
}

func (svc *Service) GetSupplier(inventory.GetSupplierRequest, *inventory.GetSupplierResponse) error {
	panic("not implemented")
}

func (svc *Service) CreateSupplier(inventory.CreateSupplierRequest, *inventory.CreateSupplierResponse) error {
	panic("not implemented")
}

func (svc *Service) UpdateSupplier(inventory.UpdateSupplierRequest, *inventory.UpdateSupplierResponse) error {
	panic("not implemented")
}

func (svc *Service) DeleteSupplier(inventory.DeleteSupplierRequest, *inventory.DeleteSupplierResponse) error {
	panic("not implemented")
}
