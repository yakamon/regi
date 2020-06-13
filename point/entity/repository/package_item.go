package repository

import "regi/point/entity"

type PackageItemRepository struct {
	storage map[string]*entity.PackageItem
}

func NewPackageItemRepository() *PackageItemRepository {
	return &PackageItemRepository{map[string]*entity.PackageItem{}}
}

func (pir *PackageItemRepository) Get(id string) (*entity.PackageItem, bool) {
	item, exists := pir.storage[id]
	return item, exists
}

func (pir *PackageItemRepository) Add(item *entity.PackageItem) {
	pir.storage[item.ID] = item
}
