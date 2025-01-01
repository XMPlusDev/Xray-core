package tcp

import (
	"github.com/xmplusdev/xray-core/v25/common"
	"github.com/xmplusdev/xray-core/v25/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
