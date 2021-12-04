package model

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CREDIT = iota + 1
	DEBT_CARD
	BANK_CARD
	VIRTUAL
)

type Collection mongo.Collection

// One of accounts' cards
type Card struct {
	Id       string   `json:"_id"`
	Owner    string   `json:"owner_id"`
	Iban     string   `json:"iban"`
	TrsctIds []string `json:"transaction_ids"`
	cache    *BalanceCache
	pin      uint16
	CardType uint8 `json:"card_type"`
}

// A struct that is used to cache the net balance until a set amount of time, so that we wouldn't
// have to calculate each transaction from the bottom
type BalanceCache struct {
	Until   time.Time
	Balance int64
}

// A struct that defines a
type Transaction struct {
	Id     string `json:"_id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}
