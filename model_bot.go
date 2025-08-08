package gamemodel

import "context"

type BotAI interface {
	Initialize(player *Player, roomInfo *RoomInfo, gameActionCh func(raw []byte, ctx context.Context) bool)
	ReceiveGameUpdate(message []byte, withSleep bool)
	Shutdown()
}

type BotFactory func() BotAI
