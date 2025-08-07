package gamemodel

type BotAI interface {
	Initialize(player *Player, roomInfo *RoomInfo, gameActionCh chan<- []byte)
	ReceiveGameUpdate(message []byte, withSleep bool)
	Shutdown()
}
