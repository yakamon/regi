package repository

import "regi/entity"

type PointCardRepository struct {
	storage map[string]*entity.PointCard
}

func NewPointCardRepository() *PointCardRepository {
	return &PointCardRepository{map[string]*entity.PointCard{}}
}

func (pir *PointCardRepository) Get(id string) (*entity.PointCard, bool) {
	card, exists := pir.storage[id]
	return card, exists
}

func (pir *PointCardRepository) Has(id string) bool {
	_, exists := pir.storage[id]
	return exists
}

func (pir *PointCardRepository) Add(card *entity.PointCard) {
	pir.storage[card.ID] = card
}
