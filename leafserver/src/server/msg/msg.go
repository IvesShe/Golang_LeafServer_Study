package msg

import "github.com/name5566/leaf/network"

// 使用默認的JSON消息處理器(另外還提供了protobuf的消息處理器)
var Processor network.Processor

func init() {
	// 注冊一個JSON消息Hello
	Processor.Register(&Hello{})
}

// 定義Hello JSON消息的結構體
type Hello struct {
	Name string
}
