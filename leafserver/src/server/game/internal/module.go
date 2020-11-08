package internal

import (
	"server/base"
	"time"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

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

func (m *Module) OnDestroy() {

}
