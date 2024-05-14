package transport

import (
	"go-projects/inventory"
	"net"
)

type InventoryTransporter interface {
	inventory.Service
	Serve(net.Listener) error
}
