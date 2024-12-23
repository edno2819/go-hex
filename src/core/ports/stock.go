package core_ports

import (
	"context"

	core_entities "github.com/edno2819/go-hex/src/core/entities"
)

type StockRepository interface {
	GetByID(ctx context.Context, id string) (*core_entities.Stock, error)
	Create(ctx context.Context, Stock *core_entities.Stock) error
	Update(ctx context.Context, Stock *core_entities.Stock) error
	Delete(ctx context.Context, Stock *core_entities.Stock) error
}
