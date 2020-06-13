package entity

import "fmt"

type Account struct {
	totalPrice float64

	usePointCard bool
	pointCardID  string

	hasInvalidCode  bool
	hasInvalidPrice bool
}

func (a *Account) Init() {
	a.totalPrice = 0
	a.usePointCard = false
	a.pointCardID = ""
	a.hasInvalidCode = false
	a.hasInvalidPrice = false
}

func (a *Account) InitWithPointCard(cardID string) {
	a.totalPrice = 0
	a.usePointCard = true
	a.pointCardID = cardID
	a.hasInvalidCode = false
	a.hasInvalidPrice = false
}

func (a *Account) UsePointCard() bool {
	return a.usePointCard
}

func (a *Account) PointCardID() string {
	return a.pointCardID
}

func (a *Account) TotalPrice() float64 {
	return a.totalPrice
}

func (a *Account) SetTotalPrice(price float64) {
	a.totalPrice = price
}

func (a *Account) AddTotalPrice(price float64) {
	a.totalPrice += price
}

func (a *Account) HasInvalidCode() bool {
	return a.hasInvalidCode
}

func (a *Account) SetHasInvalidCode(flag bool) {
	a.hasInvalidCode = flag
}

func (a *Account) HasInvalidPrice() bool {
	return a.hasInvalidPrice
}

func (a *Account) SetHasInvalidPrice(flag bool) {
	a.hasInvalidPrice = flag
}

func (a *Account) PrintAccountResult() {
	if a.HasInvalidCode() || a.HasInvalidPrice() {
		staffCall := "staff call:"
		if a.HasInvalidPrice() {
			staffCall += " 1"
		}
		if a.HasInvalidCode() {
			staffCall += " 2"
		}
		fmt.Println(staffCall)
	} else {
		fmt.Println(a.TotalPrice())
	}
}
