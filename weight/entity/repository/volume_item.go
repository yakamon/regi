package repository

import "regi/weight/entity"

type VolumeItemRepository struct {
	data map[string]*entity.VolumeItem
}

func NewVolumeItemRepository() *VolumeItemRepository {
	return &VolumeItemRepository{map[string]*entity.VolumeItem{}}
}

func (vir *VolumeItemRepository) Add(item *entity.VolumeItem) {
	vir.data[item.ID] = item
}

func (vir *VolumeItemRepository) Get(id string) (*entity.VolumeItem, bool) {
	item, ok := vir.data[id]
	return item, ok
}
