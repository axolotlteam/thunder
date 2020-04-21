package nats

import (
	"time"

	"github.com/axolotlteam/thunder/logger"
	"github.com/axolotlteam/thunder/st"
	nats "github.com/nats-io/nats.go"
)

var (
	n       *nats.Conn
	retry   = 5
	timeout = 10 * time.Second
)

// N -
func N() *nats.Conn {
	return n
}

// Connect -
func Connect(name, host string) error {

	nc, err := nats.Connect(
		host,
		nats.Name(name),                   // 連線名稱
		nats.Timeout(10*time.Second),      // 連線timeout時間
		nats.MaxReconnects(3),             // 最大重複連線次數
		nats.ReconnectWait(2*time.Second), // 每次重複連線等待時間
		nats.DontRandomize(),              // 不隨機分配
		nats.PingInterval(60*time.Second), // 每次監測間格
		nats.MaxPingsOutstanding(3),       // 幾次監測不成功後失敗
		nats.NoEcho(),                     // 是否要取得送出去的Request
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logger.WithError(err).Error("nats disconnected")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Infof("Nats reconnected to %v", nc.ConnectedUrl())

		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			logger.Infof("Nats connect closed. %v", nc.LastError())
		}),
		nats.ErrorHandler(func(nc *nats.Conn, sub *nats.Subscription, err error) {
			logger.WithError(err).Errorf("IP: %s , Sub: %v", nc.ConnectedAddr(), sub.Subject)
		}),
	)
	if err != nil {
		logger.WithError(err).Errorf("nats connect failed")
		return st.ErrorConnectFailed
	}

	n = nc

	return nil
}

// Close -
func Close() {
	if err := n.Drain(); err != nil {
		logger.WithError(err).Error("nats drain failed")
	}
}
