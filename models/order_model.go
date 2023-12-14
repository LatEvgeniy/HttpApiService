package models

type OrderModel struct {
	UserId         string  `json:"userId"`
	OrderId        string  `json:"orderId"`
	Pair           string  `json:"pair"`
	Direction      string  `json:"direction"`
	Type           string  `json:"type"`
	InitVolume     float64 `json:"initVolume"`
	InitPriice     float64 `json:"initPrice"`
	FilledPrice    float64 `json:"filledPrice"`
	FilledVolume   float64 `json:"filledVolume"`
	LockedVolume   float64 `json:"lockedVolume"`
	ExpirationDate string  `json:"expirationDate"`
	CreationDate   string  `json:"creationDate"`
	UpdatedDate    string  `json:"updatedDate"`
}
