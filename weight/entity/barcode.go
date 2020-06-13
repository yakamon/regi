package entity

import "strconv"

// ------------------------------------------------------------
// ItemType

type ItemType int

const (
	ItemTypePackage ItemType = iota
	ItemTypeVolume
)


// ------------------------------------------------------------
// BarCode

type BarCode string

func (bc BarCode) ItemType() ItemType {
	typeIdent := bc[0:2]
	switch typeIdent {
	case "02":
		return ItemTypeVolume
	default:
		return ItemTypePackage
	}
}

func (bc BarCode) IsValidCheckSum() bool {
	checkSum, _ := strconv.Atoi(string(bc[12]))
	digitSum := 0
	for _, c := range bc[0:12] {
		n, _ := strconv.Atoi(string(c))
		digitSum += n
	}
	return digitSum%10 == checkSum
}

func (bc BarCode) PackageItemID() string {
	return string(bc[0:12])
}

func (bc BarCode) VolumeItemPrice() int {
	p, _ := strconv.Atoi(string(bc[7:12]))
	return p
}

func (bc BarCode) VolumeItemID() string {
	return string(bc[2:7])
}
