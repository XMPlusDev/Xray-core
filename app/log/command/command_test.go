package command_test

import (
	"context"
	"testing"

	"github.com/xmplusdev/xray-core/v25/app/dispatcher"
	"github.com/xmplusdev/xray-core/v25/app/log"
	. "github.com/xmplusdev/xray-core/v25/app/log/command"
	"github.com/xmplusdev/xray-core/v25/app/proxyman"
	_ "github.com/xmplusdev/xray-core/v25/app/proxyman/inbound"
	_ "github.com/xmplusdev/xray-core/v25/app/proxyman/outbound"
	"github.com/xmplusdev/xray-core/v25/common"
	"github.com/xmplusdev/xray-core/v25/common/serial"
	"github.com/xmplusdev/xray-core/v25/core"
)

func TestLoggerRestart(t *testing.T) {
	v, err := core.New(&core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&log.Config{}),
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	})
	common.Must(err)
	common.Must(v.Start())

	server := &LoggerServer{
		V: v,
	}
	common.Must2(server.RestartLogger(context.Background(), &RestartLoggerRequest{}))
}
