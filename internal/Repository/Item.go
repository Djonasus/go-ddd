package repository

import entity "ddd/internal/Entity"

type ItemRepository interface {
	// Crud
	Create(*entity.Item) error
	Remove(*entity.Item) error
	Update(*entity.Item) error
	GetById(*entity.Item) error
}
