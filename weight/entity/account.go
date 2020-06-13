package entity

import "fmt"

type Account struct {
	TotalPrice int

	PrevBarCode      BarCode
	PrevBucketWeight float64

	HasInvalidCode   bool
	HasInvalidWeight bool
}

func (a *Account) Init() {
	a.TotalPrice = 0
	a.PrevBarCode = ""
	a.PrevBucketWeight = 0
	a.HasInvalidCode = false
	a.HasInvalidWeight = false
}

func (a *Account) PrintAccountResult() {
	if a.HasInvalidCode || a.HasInvalidWeight {
		staffCall := "staff call:"
		if a.HasInvalidCode {
			staffCall += " 1"
		}
		if a.HasInvalidWeight {
			staffCall += " 2"
		}
		fmt.Println(staffCall)
	} else {
		fmt.Println(a.TotalPrice)
	}
}
