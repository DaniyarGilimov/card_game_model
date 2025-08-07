package gamemodel

import model "github.com/daniyargilimov/card_api_model"

// PlayerInstruction plm
type PlayerInstruction struct {
	ID         int                    `json:"id"`
	Username   string                 `json:"username"`
	Avatar     string                 `json:"avatar"`
	TablePlace int                    `json:"tablePlace"`
	Inventory  *model.PersonInventory `json:"inventory"`
}

// PlayerJoinedInstruction is used to convert to json
type PlayerJoinedInstruction struct {
	Instruction string `json:"status"`
	Data        struct {
		ID         int                    `json:"id"`
		Username   string                 `json:"username"`
		Avatar     string                 `json:"avatar"`
		TablePlace int                    `json:"tablePlace"`
		Inventory  *model.PersonInventory `json:"inventory"`
	} `json:"data"`
	Order int `json:"order"`
}

// PlayerLeftInstruction is used to convert to json
type PlayerLeftInstruction struct {
	Instruction string `json:"status"`
	Data        struct {
		ID int `json:"playerId"`
	} `json:"data"`
}

type PlayerNotEnoughChipsInstruction struct {
	Instruction string             `json:"status"`
	Data        NotEnoughMoneyData `json:"data"`
}

type UtilUpdateInventory struct {
	Instruction string               `json:"status"`
	Data        *UpdateInventoryData `json:"data"`
}

type UpdateInventoryData struct {
	PlayerID  int                    `json:"playerId"`
	Inventory *model.PersonInventory `json:"inventory"`
}

type NotEnoughMoneyData struct {
	ID          int   `json:"playerId"`
	ChipsAmount int64 `json:"chipsAmount"`
}

// GameBridgeInstructer is used to convert to json
type GameBridgeInstructer struct {
	Instruction string `json:"status"`
	Data        struct {
		ID        int    `json:"playerId"`
		HashOrder string `json:"hashOrder"`
	} `json:"data"`
}

// GeneralInstructer is used to convert to json
type GeneralInstructer struct {
	Instruction string      `json:"status"`
	Data        interface{} `json:"data"`
}

// RoomCreatedInstruction is used to send room id for creator
type RoomCreatedInstruction struct {
	Instruction string `json:"status"`
	Data        struct {
		RoomID      int                      `json:"roomId"`
		OwnData     *model.PersonInventory   `json:"ownData"`
		Inventories []*model.PersonInventory `json:"inventories"`
		RoomName    string                   `json:"roomName"`
		Messages    []*model.SingleMessage   `json:"messages"`
	} `json:"data"`
	Order int `json:"order"`
	ID    int `json:"id"`
}

// ConnectionCaseOneInstruction is when room has 1 player
type ConnectionCaseOneInstruction struct {
	Instruction string `json:"status"`
	Data        struct {
		Players      []PlayerInstruction    `json:"players"`
		OwnData      *model.PersonInventory `json:"ownData"`
		Inventory    *model.PersonInventory `json:"inventory"`
		BookedPlaces []int                  `json:"bookedPlaces"`
		RoomName     string                 `json:"roomName"`
		Messages     []*model.SingleMessage `json:"messages"`
	} `json:"data"`
	Order int `json:"order"`
	ID    int `json:"id"`
}

// ConnectionCaseJoinInstruction
type ConnectionCaseJoinInstruction struct {
	Instruction string `json:"status"`
	Data        struct {
		BetStates []*PlayerBetStateResponse `json:"betStates"`
		CardsData struct {
			Trump string           `json:"trump"`
			Cards []*PlayerRequest `json:"cards"`
		} `json:"cardsData"`

		LoadingPlayerIDs []int               `json:"playerIds"`
		OtherPlayers     []PlayerInstruction `json:"otherPlayers"`

		Inventory *model.PersonInventory `json:"inventory"`

		BookedPlaces []int                  `json:"bookedPlaces"`
		BankValue    int64                  `json:"bankValue"`
		RoomName     string                 `json:"roomName"`
		Messages     []*model.SingleMessage `json:"messages"`
	} `json:"data"`
	Order int `json:"order"`
	ID    int `json:"id"`
}

// ConnectionCaseNeedInstruction
type ConnectionCaseNeedInstruction struct {
	Instruction string `json:"status"`
	Data        struct {
		BetStates []*PlayerBetStateResponse `json:"betStates"`
		CardsData struct {
			Trump string           `json:"trump"`
			Cards []*PlayerRequest `json:"cards"`
		} `json:"cardsData"`

		CurrentNeed LoadingTurnJoinRound `json:"currentNeed"`

		OtherPlayers []PlayerInstruction `json:"otherPlayers"`

		Inventory *model.PersonInventory `json:"inventory"`

		BookedPlaces []int                  `json:"bookedPlaces"`
		BankValue    int64                  `json:"bankValue"`
		RoomName     string                 `json:"roomName"`
		Messages     []*model.SingleMessage `json:"messages"`
	} `json:"data"`
	Order int `json:"order"`
	ID    int `json:"id"`
}

// ConnectionCasePlayInstruction
type ConnectionCasePlayInstruction struct {
	Instruction string `json:"status"`
	Data        struct {
		CardsData struct {
			Trump        string           `json:"trump"`
			Cards        []*PlayerRequest `json:"cards"`
			CentralCards []*MovedCard     `json:"centralCards"`
		} `json:"cardsData"`

		CurrentPlay LoadingTurnPlayInstruction `json:"currentPlay"`

		OtherPlayers []PlayerInstruction       `json:"otherPlayers"`
		BetStates    []*PlayerBetStateResponse `json:"betStates"`

		Inventory *model.PersonInventory `json:"inventory"`

		BookedPlaces []int                  `json:"bookedPlaces"`
		BankValue    int64                  `json:"bankValue"`
		RoomName     string                 `json:"roomName"`
		Messages     []*model.SingleMessage `json:"messages"`
	} `json:"data"`
	Order int `json:"order"`
	ID    int `json:"id"`
}

// ExportRoomData for exporting roomData
type ExportRoomData struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	PlayersCount int    `json:"playersCount"`
	RoomSize     int    `json:"roomSize"`
	InitialBet   int64  `json:"initialBet"`
	IsOpen       bool   `json:"isOpen"`
	Password     string `json:"password"`
}

// ExportAllRoomsData for exporting all rooms
type ExportAllRoomsData struct {
	InitialBet    int64    `json:"initialBet"`
	PlayersAmount int      `json:"playersAmount"`
	PlaceName     string   `json:"placeName"`
	Wallpaper     string   `json:"wallpaper"`
	PlayersImage  []string `json:"playersImage"`
	IsPrivate     bool     `json:"isPrivate"`
}

// ExportAllRoomsBotData for exporting all rooms
type ExportAllRoomsBotData struct {
	RoomName     string `json:"roomName"`
	RoomID       int    `json:"roomID"`
	InitialBet   int    `json:"initialBet"`
	PlayersCount int    `json:"playersCount"`
}

// ExportData for general purpose
type ExportData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// CardDistributionInstruction used to send distribution
type CardDistributionInstruction struct {
	Status string `json:"status"`
	Data   struct {
		Cards []*PlayerRequest `json:"cards"`
	} `json:"data"`
	Order int `json:"order"`
}

// LoadingTurnBetInstruction used to set turn
type LoadingTurnBetInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID                  int     `json:"playerId"`
		CallValue                 int64   `json:"callValue"`
		RaiseRange                []int64 `json:"raiseRange"`
		HashOrder                 string  `json:"hashOrder"` //Randomly generated hash order
		IsResetPlayersBetStatuses bool    `json:"isResetPlayersBetStatuses"`
		// Svara case
	} `json:"data"`
	Order int `json:"order"`
}

type LoadingTurnPlayInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID int   `json:"playerId"`
		CardIDs  []int `json:"cardIds"`
	} `json:"data"`
}

type LoadingTurnJoinRound struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID int `json:"playerId"`
	} `json:"data"`
}

// // LoadingSvaraInstruction used to set svara
// type LoadingSvaraInstruction struct {
// 	Status string `json:"status"`
// 	Data   struct {
// 		UsedCombinations []*UsedCombination `json:"usedCombinations"`
// 		CallValue        int                `json:"callValue"`
// 		SvaraPlayerIDs   []int              `json:"svaraPlayerIds"`
// 		PlayerIDs        []int              `json:"playerIds"`
// 		PlayersShowData  []*PlayerShowData  `json:"playersShowData"`
// 	} `json:"data"`
// 	Order int `json:"order"`
// }

// UtilCardsOpenInstruction used to open cards
type UtilCardsOpenInstruction struct {
	Status string `json:"status"`
	Data   struct {
	} `json:"data"`
	Order int `json:"order"`
}

// MoveTurnBetInstruction struct used to set move turn
type MoveTurnBetInstruction struct {
	Status string `json:"status"`
	Data   struct {
		Amount    int64  `json:"amount"`
		PlayerID  int    `json:"playerId"`
		BetStatus string `json:"betStatus"`
	} `json:"data"`
	Order int `json:"order"`
}

// MoveTurnCardsNeededInstruction struct used to set move turn
type MoveTurnCardsNeededInstruction struct {
	Status string `json:"status"`
	Data   struct {
		// Player sends to backend
		PlayerID       int   `json:"playerId"`
		RemovedCardIds []int `json:"cardIds"`

		// Backend generates these values and broadcasts
		Amount       int64   `json:"amount"`
		AddedCards   []*Card `json:"addedCards"`
		RemovedCards []*Card `json:"removedCards"`
	} `json:"data"`
	Order int `json:"order"`
}

// UtilTotalBetInstruction to send that betting is ended
type UtilTotalBetInstruction struct {
	Status string `json:"status"`
	Data   struct {
		TotalChips  int64                    `json:"totalChips"`
		PlayerIDs   []int                    `json:"playerIds"`
		Inventories []*model.PersonInventory `json:"inventories"`
	} `json:"data"`
	Order int `json:"order"`
}

// GameWinInstruction to send that betting is ended
type GameWinInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID    int                      `json:"playerId"`
		Inventories []*model.PersonInventory `json:"inventories"`
		WonAmount   int64                    `json:"wonAmount"`
	} `json:"data"`
	Order int `json:"order"`
}

// MoveTurnPlayInstruction used to move card
type MoveTurnPlayInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID int `json:"playerId"`
		CardID   int `json:"cardId"`
	} `json:"data"`
}

// RoundWinInstruction used to send win instruction
type RoundWinInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID      int `json:"playerId"`
		WonRoundCount int `json:"wonRoundCount"`
	} `json:"data"`
}

// UsedCombination gives information od cards combination that was used
type UsedCombination struct {
	PlayerID int   `json:"playerId"`
	CardIDs  []int `json:"cardIds"`
	Point    int   `json:"point"` //Total points
}

// ShowMessage is used to send message
type ShowMessage struct {
	Status string `json:"status"`
	Data   struct {
		Message          string             `json:"message"`
		Status           string             `json:"status"`
		UsedCombinations []*UsedCombination `json:"usedCombinations"`
	} `json:"data"`
	Order int `json:"order"`
}

// StatusInstruction status
type StatusInstruction struct {
	Status string `json:"status"`
	Data   struct {
		InitialBet int64 `json:"initialBet"`
	} `json:"data"`
}

// EmojiInstruction EmojiInstruction
type EmojiInstruction struct {
	Status string `json:"status"`
	Data   struct {
		EmojiID  int `json:"emojiId"`
		PlayerID int `json:"playerId"`
	} `json:"data"`
}

// RecieveThrowObjectInstruction is used to throw objects
type RecieveThrowObjectInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID   int `json:"playerId"`
		ThrowObjID int `json:"throwObjId"`
	} `json:"data"`
}

// InviteFriendInstruction is used to invite friend
type InviteFriendInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID   int `json:"playerId"`
		TablePlace int `json:"tablePlace"`
	} `json:"data"`
}

// CancleInvitationInstruction is used to cancle invitation
type CancleInvitationInstruction struct {
	Status string `json:"status"`
	Data   struct {
		PlayerID int `json:"playerId"`
		RoomID   int `json:"roomId"`
	} `json:"data"`
}

// BookPlaceInstruction is used to book one place for friend
type BookPlaceInstruction struct {
	Status string `json:"status"`
	Data   struct {
		TablePlace int `json:"tablePlace"`
	} `json:"data"`
}

// InvitationInstruction is used to send invitation
type InvitationInstruction struct {
	Status string `json:"status"`
	Data   struct {
		DisplayName string `json:"displayName" bson:"displayName"`
		Avatar      string `json:"photoUrl" bson:"photoUrl"`
		InitialBet  int    `json:"initialBet" bson:"initialBet"`
		RoomID      int    `json:"roomId" bson:"roomId"`
		PlayerID    int    `json:"playerId" bson:"playerId"`
	} `json:"data"`
}

// SendThrowObjectInstruction is used to get
type SendThrowObjectInstruction struct {
	Status string         `json:"status"`
	Data   *SendThrowData `json:"data"`
}

// SendThrowData is used to get
type SendThrowData struct {
	ThrowObjID   int                         `json:"throwObjId"`
	ToPlayerID   int                         `json:"toPlayerId"`
	FromPlayerID int                         `json:"fromPlayerId"`
	Inventory    []*model.ThrowableInventory `json:"inventory"`
}
