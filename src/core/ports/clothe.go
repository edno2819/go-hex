package core_ports

import core_entities "github.com/edno2819/go-hex/src/core/entities"

type ClotheRepository interface {
	GetByID(id string) (*core_entities.Clothe, error)
	Create(clothe *core_entities.Clothe) error
	Update(clothe *core_entities.Clothe) error
	Delete(id string) error
}
