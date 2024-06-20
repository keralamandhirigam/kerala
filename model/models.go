package model

import (
	"time"
)

type CreateWitchcraft struct {
	WhichcraftId     string    `json:"witchcraftId"`
	Name             string    `json:"name"`
	Address          string    `json:"address"`
	Place            string    `json:"place"`
	Remarks          string    `json:"remarks"`
	Payment          int       `json:"payment"`
	BalanceAmount    float64   `json:"balanceAmount"`
	VisitorsDuration time.Time `json:"visitorsDuration"`
}

type UpdateWitchcraft struct {
	Name             string    `json:"name"`
	Address          string    `json:"address"`
	Place            string    `json:"place"`
	Remarks          string    `json:"remarks"`
	Payment          int       `json:"payment"`
	BalanceAmount    float64   `json:"balanceAmount"`
	VisitorsDuration time.Time `json:"VisitorsDuration"`
}

type Witchcraft struct {
	WitchcraftId     string    `json:"witchcraftId"`
	Name             string    `json:"name"`
	Address          string    `json:"address"`
	Place            string    `json:"place"`
	Remarks          string    `json:"remarks"`
	Payment          int       `json:"payment"`
	BalanceAmount    float64   `json:"balanceAmount"`
	VisitorsDuration time.Time `json:"VisitorsDuration"`
}

type WitchcraftDAO struct {
	WitchcraftId     string    `json:"witchcraftId"`
	Name             string    `bson:"name"`
	Address          string    `bson:"address"`
	Place            string    `bson:"place"`
	Remarks          string    `bson:"remarks"`
	Payment          int       `bson:"payment"`
	BalanceAmount    float64   `bson:"balanceAmount"`
	VisitorsDuration time.Time `bson:"VisitorsDuration"`
}
