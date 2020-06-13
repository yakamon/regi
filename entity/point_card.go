package entity

type PointCard struct {
	ID    string
	point int
}

func NewPointCard(id string) *PointCard {
	return &PointCard{
		ID:    id,
		point: 0,
	}
}

func (pc *PointCard) Point() int {
	return pc.point
}

func (pc *PointCard) SetPoint(point int) {
	pc.point = point
}
