package core_ports

import (
	"context"

	core_entities "github.com/edno2819/go-hex/src/core/entities"
)

type ClotheRepository interface {
	GetByID(ctx context.Context, id string) (*core_entities.Clothe, error)
	List(ctx context.Context, skip, limit uint64) ([]core_entities.Clothe, error)
	Create(ctx context.Context, clothe *core_entities.Clothe) error
	Update(ctx context.Context, clothe *core_entities.Clothe) error
	Delete(ctx context.Context, id string) error
}
