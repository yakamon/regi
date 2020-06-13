package repository

import "regi/entity"

type VolumeItemRepository struct {
	storage map[string]*entity.VolumeItem
}

func NewVolumeItemRepository() *VolumeItemRepository {
	return &VolumeItemRepository{map[string]*entity.VolumeItem{}}
}

func (pir *VolumeItemRepository) Get(id string) (*entity.VolumeItem, bool) {
	item, exists := pir.storage[id]
	return item, exists
}

func (pir *VolumeItemRepository) Add(item *entity.VolumeItem) {
	pir.storage[item.ID] = item
}
