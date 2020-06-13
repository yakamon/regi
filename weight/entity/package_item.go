package entity

import "math"

type PackageItem struct {
	ID                   string
	Price                int
	StandardWeight       float64
	AllowableErrorWeight float64
}

func (pi *PackageItem) IsValidWeightPackageItem(weight float64) bool {
	return math.Abs(pi.StandardWeight-weight) <= pi.AllowableErrorWeight
}
