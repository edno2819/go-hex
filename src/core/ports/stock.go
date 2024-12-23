package core_ports

import core_entities "github.com/edno2819/go-hex/src/core/entities"

type StockRepository interface {
	GetByID(id string) (*core_entities.Stock, error)
	Create(Stock *core_entities.Stock) error
	Update(Stock *core_entities.Stock) error
	Delete(Stock *core_entities.Stock) error
}
