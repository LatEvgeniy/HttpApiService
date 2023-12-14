package models

type UserModel struct {
	Id       string          `json:"id"`
	Name     string          `json:"name"`
	Balances []*BalanceModel `json:"balances"`
}

type BalanceModel struct {
	CurrencyName  string  `json:"currencyName"`
	Balance       float64 `json:"balance" jsondefault:"0.0"`
	LockedBalance float64 `json:"lockedBalance" jsondefault:"0.0"`
}
