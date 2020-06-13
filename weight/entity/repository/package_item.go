package repository

import "regi/weight/entity"

type PackageItemRepository struct {
	data map[string]*entity.PackageItem
}

func NewPackageItemRepository() *PackageItemRepository {
	return &PackageItemRepository{map[string]*entity.PackageItem{}}
}

func (pir *PackageItemRepository) Add(item *entity.PackageItem) {
	pir.data[item.ID] = item
}

func (pir *PackageItemRepository) Get(id string) (*entity.PackageItem, bool) {
	item, ok := pir.data[id]
	return item, ok
}
