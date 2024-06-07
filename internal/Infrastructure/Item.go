package infrastructure

import (
	entity "ddd/internal/Entity"
	repository "ddd/internal/Repository"

	"gorm.io/gorm"
)

type ItemInfrastructure struct {
	Db *gorm.DB
}

// Create implements repository.ItemRepository.
func (i *ItemInfrastructure) Create(*entity.Item) error {
	panic("unimplemented")
}

// GetById implements repository.ItemRepository.
func (i *ItemInfrastructure) GetById(*entity.Item) error {
	panic("unimplemented")
}

// Remove implements repository.ItemRepository.
func (i *ItemInfrastructure) Remove(*entity.Item) error {
	panic("unimplemented")
}

// Update implements repository.ItemRepository.
func (i *ItemInfrastructure) Update(*entity.Item) error {
	panic("unimplemented")
}

func NewItemInfrastructure(db *gorm.DB) repository.ItemRepository {
	return &ItemInfrastructure{
		Db: db,
	}
}
