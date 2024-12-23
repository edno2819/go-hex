package core_serices

import (
	"context"

	core_entities "github.com/edno2819/go-hex/src/core/entities"
	core_ports "github.com/edno2819/go-hex/src/core/ports"
)

type ClotheService struct {
	repo core_ports.ClotheRepository
}

func NewClotheService(repo core_ports.ClotheRepository) *ClotheService {
	return &ClotheService{repo}
}

func (s *ClotheService) GetByID(ctx context.Context, id string) (*core_entities.Clothe, error) {
	clothe, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return clothe, nil
}

func (s *ClotheService) List(ctx context.Context, skip uint64, limit uint64) ([]core_entities.Clothe, error) {
	clothes, err := s.repo.List(ctx, skip, limit)
	if err != nil {
		return nil, err
	}
	return clothes, nil
}

func (s *ClotheService) Create(ctx context.Context, clothe *core_entities.Clothe) error {
	err := s.repo.Create(ctx, clothe)
	return err
}

func (s *ClotheService) Update(ctx context.Context, clothe *core_entities.Clothe) error {
	err := s.repo.Create(ctx, clothe)
	return err
}

func (s *ClotheService) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	return err
}
