package entity

import "math"

type VolumeItem struct {
	ID                   string
	WeightPerHundredYen  float64
	PackageWeight        float64
	AllowableErrorWeight float64
}

func (vi *VolumeItem) StandardWeight(price int) float64 {
	return vi.WeightPerHundredYen*float64(price)/100 + vi.PackageWeight
}

func (vi *VolumeItem) IsValidWeightVolumeItem(weight float64, price int) bool {
	return math.Abs(vi.StandardWeight(price)-weight) <= vi.AllowableErrorWeight
}
