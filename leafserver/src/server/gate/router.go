package gate

import (
	"server/game"
	"server/msg"
)

func init() {
	// 指定消息Hello路由到game模塊
	// 模塊間使用ChanRPC通訊，消息路由也不例外
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
