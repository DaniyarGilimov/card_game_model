package gamemodel

import (
	model "github.com/daniyargilimov/card_api_model"
)

// Player tosibosi
type Player struct {
	Name      string `json:"-"`
	Token     string `json:"-"`
	PushToken string `json:"-"`
	PlayerID  int    `json:"playerId"`
	IsBot     bool   `json:"isBot,omitempty"` // True if this player is a bot

	Inventory   *model.PersonInventory
	RuntimeData *PlayerRuntimeData
	// Cards      []*Card `json:"cards"`
	// MovedCards []*Card `json:"-"`

	// WonCards    []*MovedCard    `json:"-"`
	// PlayerState *PlayerState   `json:"-"`
	// OwnData     *model.OwnData `json:"ownData"`

}

// PlayerBetStateResponse model for communicating with frontend
type PlayerBetStateResponse struct {
	PlayerID  int    `json:"playerId"`
	BetAmount int64  `json:"betAmount"`
	BetStatus string `json:"betStatus"`
}

func (p *Player) ToPlayerRequest() *PlayerRequest {
	cards := p.RuntimeData.Cards
	if cards == nil {
		cards = make([]*Card, 0)
	}

	return &PlayerRequest{
		PlayerID: p.PlayerID,
		Cards:    cards,
	}
}

// PlayerRequest model for communicating with frontend
type PlayerRequest struct {
	PlayerID int     `json:"playerId"`
	Cards    []*Card `json:"cards"`
}

type PlayerRuntimeData struct {
	HasEnoughChips bool `json:"-"`
	// BetAmount Used to display for frontend, it resets in draw phase
	BetAmount int64 `json:"betAmount"`

	CardsToExchangeCount int
	// TotalBetAmount shows actual bet amount of a player, both in phase 1 and draw phase
	TotalBetAmount int64   `json:"totalBetAmount"`
	BetStatus      string  `json:"betStatus"`
	AfkCount       int     `json:"-"`
	Cards          []*Card `json:"-"`
}

// Clear is used to reset all values
func (runtimeData *PlayerRuntimeData) Clear() {
	runtimeData.HasEnoughChips = true
	runtimeData.BetAmount = 0
	runtimeData.TotalBetAmount = 0
	runtimeData.BetStatus = ""
}

// PlayerBetState betstate
type PlayerBetState struct {
	PlayerID       int    `json:"playerId"`
	BetAmount      int    `json:"betAmount"`
	BetStatus      string `json:"betStatus"`
	BetAmountLast  int    `json:"-"` //Used in svara case to detect deltaBet
	BetStatusLocal string `json:"-"`
	MadeRaiseCount int    `json:"-"`
}

type MovedCard struct {
	PlayerID int   `json:"playerId"`
	Card     *Card `json:"card"`
}

// Card is game card
type Card struct {
	CardID int    `json:"cardId"` //0-35
	Value  int    `json:"value"`  //2-14
	Suit   string `json:"suit"`   //CROSS SPADES HEART DIMOND

	Importance int `json:"-"`
}
