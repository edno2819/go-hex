package core_ports

import (
	"context"

	core_entities "github.com/edno2819/go-hex/src/core/entities"
)

type OrderRepository interface {
	GetByID(ctx context.Context, id string) (*core_entities.Order, error)
	Create(ctx context.Context, Order *core_entities.Order) error
	Update(ctx context.Context, Order *core_entities.Order) error
	Delete(ctx context.Context, Order *core_entities.Order) error
}
