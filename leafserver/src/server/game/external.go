package game

import (
	"server/game/internal"
)

var (
	// 實例化game模塊
	Module = new(internal.Module)
	// 暴露ChanRPC
	ChanRPC = internal.ChanRPC
)
