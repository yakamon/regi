package main

import (
	"bufio"
	"os"
	"regi/weight/entity"
	"regi/weight/entity/repository"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)

	account               = new(entity.Account)
	packageItemRepository = repository.NewPackageItemRepository()
	volumeItemRepository  = repository.NewVolumeItemRepository()
)

func main() {
	PrepareProductRepositories()
	for scanner.Scan() {
		HandleAccount(scanner.Text())
	}
}

func PrepareProductRepositories() {
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < N; i++ {
		scanner.Scan()
		itemInfo := strings.Split(scanner.Text(), " ")

		switch itemID := itemInfo[0]; len(itemID) {
		case 12:
			price, _ := strconv.Atoi(itemInfo[1])
			standardWeight, _ := strconv.ParseFloat(itemInfo[2], 64)
			allowableErrorWeight, _ := strconv.ParseFloat(itemInfo[3], 64)

			packageItemRepository.Add(&entity.PackageItem{
				ID:                   itemID,
				Price:                price,
				StandardWeight:       standardWeight,
				AllowableErrorWeight: allowableErrorWeight,
			})
		case 5:
			weightPerHundredYen, _ := strconv.ParseFloat(itemInfo[1], 64)
			packageWeight, _ := strconv.ParseFloat(itemInfo[2], 64)
			allowableErrorWeight, _ := strconv.ParseFloat(itemInfo[3], 64)

			volumeItemRepository.Add(&entity.VolumeItem{
				ID:                   itemID,
				WeightPerHundredYen:  weightPerHundredYen,
				PackageWeight:        packageWeight,
				AllowableErrorWeight: allowableErrorWeight,
			})
		}
	}
}

func HandleAccount(line string) {
	switch {
	case strings.HasPrefix(line, "start"): // Start accounting
		account.Init()
	case strings.HasPrefix(line, "end"): // Complete accounting
		scanInfo := strings.Split(line, " ")
		curBucketWeight, _ := strconv.ParseFloat(scanInfo[1], 64)

		// Check previous item
		if account.PrevBarCode != "" {
			handleAccountItem(curBucketWeight)
		}

		// Console output result
		account.PrintAccountResult()
	default:
		scanInfo := strings.Split(line, " ")
		barCode := scanInfo[0]
		curBucketWeight, _ := strconv.ParseFloat(scanInfo[1], 64)

		if account.PrevBarCode != "" {
			handleAccountItem(curBucketWeight)
		}

		account.PrevBarCode = entity.BarCode(barCode)
		account.PrevBucketWeight = curBucketWeight
	}
}

func handleAccountItem(curBucketWeight float64) {
	if !account.PrevBarCode.IsValidCheckSum() {
		account.HasInvalidCode = true
		return
	}
	prevItemWeight := curBucketWeight - account.PrevBucketWeight
	switch account.PrevBarCode.ItemType() {
	case entity.ItemTypePackage:
		handleAccountPackageItem(prevItemWeight)
	case entity.ItemTypeVolume:
		handleAccountVolumeItem(prevItemWeight)
	}
}

func handleAccountPackageItem(prevItemWeight float64) {
	prevItem, exists := packageItemRepository.Get(account.PrevBarCode.PackageItemID())
	if !exists {
		account.HasInvalidCode = true
		return
	}
	if !prevItem.IsValidWeightPackageItem(prevItemWeight) {
		account.HasInvalidWeight = true
		return
	}
	account.TotalPrice += prevItem.Price
}

func handleAccountVolumeItem(prevItemWeight float64) {
	prevItem, exists := volumeItemRepository.Get(account.PrevBarCode.VolumeItemID())
	if !exists {
		account.HasInvalidCode = true
		return
	}
	if !prevItem.IsValidWeightVolumeItem(prevItemWeight, account.PrevBarCode.VolumeItemPrice()) {
		account.TotalPrice += account.PrevBarCode.VolumeItemPrice()
		return
	}
	account.HasInvalidWeight = true
}
