package core_ports

import core_entities "github.com/edno2819/go-hex/src/core/entities"

type OrderRepository interface {
	GetByID(id string) (*core_entities.Order, error)
	Create(Order *core_entities.Order) error
	Update(Order *core_entities.Order) error
	Delete(Order *core_entities.Order) error
}
