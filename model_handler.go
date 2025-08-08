package gamemodel

import (
	"context"

	model "github.com/daniyargilimov/card_api_model"
)

type GameFactory func(
	ctx context.Context,
	gameCtxCancel context.CancelFunc,
	roomInfo *RoomInfo,
	service Service,
	roomBroadcastCh chan []byte,
	roomBroadcastExceptCh chan UserIdAndByte,
	roomUnicastCh chan UserIdAndByte) Game

type Service interface {
	GetMessages() ([]*model.SingleMessage, error)
	IncrementChips(id int, chips int64) (int64, error)
}

type Game interface {
	SendRoomToGameInst(raw []byte, ctx context.Context) bool
	SendRoomToGamePlayerJoin(player *Player, ctx context.Context) bool
	SendRoomToGamePlayerLeave(player *Player, ctx context.Context) bool
	RunGameListener()
	GetPlayers() []*Player
	IsFastForward() bool
}

// RoomInfo is room's information, that is static
type RoomInfo struct {
	Name         string //`json:"title"`
	RoomSize     int    //`json:"room_size"`
	InitialBet   int64  //`json:"initial_bet"`
	IsOpen       bool   //`json:"is_open"`
	Password     string //`json: "password"`
	DefinedCards bool   // Defines if a room distributes predefined cards

	// TrumpSuit    string

	ID int
}

type UserIdAndByte struct {
	UserID int
	Byte   []byte
}
