package internal

import (
	"server/conf"
	"server/game"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/prometheus/common/log"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}

	log.Debug("1")

	// 定义变量 res 接收结果
	var res string

	skeleton.Go(func() {
		// 这里使用 Sleep 来模拟一个很慢的操作
		time.Sleep(1 * time.Second)

		// 假定得到结果
		res = "3"
	}, func() {
		log.Debug(res)
	})

	log.Debug("2")
}
