package main

import (
	"bufio"
	"math"
	"os"
	"regi/entity"
	"regi/entity/repository"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)

	account               = new(entity.Account)
	packageItemRepository = repository.NewPackageItemRepository()
	volumeItemRepository  = repository.NewVolumeItemRepository()
	pointCardRepository   = repository.NewPointCardRepository()
)

func main() {
	PrepareItemRepositories()
	for scanner.Scan() {
		HandleAccounts(scanner.Text())
	}
}

func PrepareItemRepositories() {
	scanner.Scan()
	numOfItems, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < numOfItems; i++ {
		scanner.Scan()
		itemInfo := strings.Split(scanner.Text(), " ")

		switch itemID := itemInfo[0]; len(itemID) {
		case 5: // is VolumeProduct
			volumeItemRepository.Add(&entity.VolumeItem{
				ID: itemID,
			})
		case 12: // is PackageProduct
			price, _ := strconv.ParseFloat(itemInfo[1], 64)
			packageItemRepository.Add(&entity.PackageItem{
				ID:    itemID,
				Price: price,
			})
		}
	}
}

func HandleAccounts(line string) {
	switch {
	case strings.HasPrefix(line, "start"): // start accounting
		info := strings.Split(line, " ")

		// initialize accounting
		if len(info) < 2 {
			account.Init()
		} else {
			cardID := info[1]
			pointCardRepository.Add(entity.NewPointCard(cardID))
			account.InitWithPointCard(cardID)
		}
	case strings.HasPrefix(line, "end"): // complete accounting
		if account.UsePointCard() {
			// get point card
			card, _ := pointCardRepository.Get(account.PointCardID())

			// consume points
			totalPriceAfterConsumePoint := account.TotalPrice() - float64(card.Point())
			if totalPriceAfterConsumePoint < 0 {
				card.SetPoint(int(math.Abs(totalPriceAfterConsumePoint)))
			} else {
				// give points
				card.SetPoint(int(math.Round(totalPriceAfterConsumePoint / 100)))
				// set total price
				account.SetTotalPrice(totalPriceAfterConsumePoint)
			}
		}

		// console output result
		account.PrintAccountResult()
	default:
		handleAccountItem(entity.BarCode(line))
	}
}

func handleAccountItem(barCode entity.BarCode) {
	// validate bar-code length
	if !barCode.IsValidLength() {
		account.SetHasInvalidCode(true)
		return
	}

	// validate bar-code check-sum
	if !barCode.IsValidCheckSum() {
		account.SetHasInvalidCode(true)
		return
	}

	// handle item by item type
	switch barCode.ItemType() {
	case entity.ItemTypePackage:
		handleAccountPackageItem(barCode)
	case entity.ItemTypeVolume:
		handleAccountVolumeItem(barCode)
	}
}

func handleAccountPackageItem(code entity.BarCode) {
	item, exists := packageItemRepository.Get(code.PackageProductItemID())
	if !exists {
		account.SetHasInvalidCode(true)
		return
	}

	switch {
	case code.IsBeforeDiscount():
		account.AddTotalPrice(item.Price)

	case code.IsAfterPercentDiscount():
		discount := item.Price * code.PercentDiscountRate()
		discountedPrice := item.Price - discount
		account.AddTotalPrice(math.Round(discountedPrice))

	case code.IsAfterAmountDiscount():
		discountedPrice := item.Price - code.AmountDiscountRate()
		if discountedPrice < 0 {
			account.SetHasInvalidPrice(true)
			return
		}
		account.AddTotalPrice(discountedPrice)

	}
}

func handleAccountVolumeItem(code entity.BarCode) {
	_, exists := volumeItemRepository.Get(code.VolumeProductItemID())
	if !exists {
		account.SetHasInvalidCode(true)
		return
	}

	switch {
	case code.IsBeforeDiscount():
		account.AddTotalPrice(code.VolumeProductItemPrice())

	case code.IsAfterPercentDiscount():
		discount := code.VolumeProductItemPrice() * code.PercentDiscountRate()
		discountedPrice := code.VolumeProductItemPrice() - discount
		account.AddTotalPrice(math.Round(discountedPrice))

	case code.IsAfterAmountDiscount():
		discountedPrice := code.VolumeProductItemPrice() - code.AmountDiscountRate()
		if discountedPrice < 0 {
			account.SetHasInvalidPrice(true)
			return
		}
		account.AddTotalPrice(discountedPrice)

	}
}
