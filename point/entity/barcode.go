package entity

import "strconv"

// --------------------------------------------------------------------------------
// ItemType

type ItemType int

const (
	ItemTypePackage ItemType = iota
	ItemTypeVolume
)

// --------------------------------------------------------------------------------
// BarCode

type BarCode string

func (bc BarCode) ItemType() ItemType {
	itemTypeIdent := bc[0:2]
	switch itemTypeIdent {
	case "02":
		return ItemTypeVolume
	default:
		return ItemTypePackage
	}
}

func (bc BarCode) IsValidLength() bool {
	barCodeLength := len(bc)
	return barCodeLength == 13 || barCodeLength == 15 || barCodeLength == 18
}

func (bc BarCode) IsBeforeDiscount() bool {
	return len(bc) == 13
}

func (bc BarCode) IsAfterPercentDiscount() bool {
	return len(bc) == 15
}

func (bc BarCode) IsAfterAmountDiscount() bool {
	return len(bc) == 18
}

func (bc BarCode) IsValidCheckSum() bool {
	last := len(bc) - 1
	checkSum, _ := strconv.Atoi(string(bc[last]))
	digitSum := 0
	for _, c := range bc[:last] {
		n, _ := strconv.Atoi(string(c))
		digitSum += n
	}
	return digitSum%10 == checkSum
}

func (bc BarCode) PackageProductItemID() string {
	return string(bc[:12])
}

func (bc BarCode) VolumeProductItemID() string {
	return string(bc[2:7])
}

func (bc BarCode) VolumeProductItemPrice() float64 {
	p, _ := strconv.ParseFloat(string(bc[7:12]), 64)
	return p
}

func (bc BarCode) PercentDiscountRate() float64 {
	r, _ := strconv.ParseFloat(string(bc[12:14]), 64)
	return r / 100
}

func (bc BarCode) AmountDiscountRate() float64 {
	r, _ := strconv.ParseFloat(string(bc[12:17]), 64)
	return r
}
