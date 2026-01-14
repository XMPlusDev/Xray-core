package udp

import (
	"github.com/xmplusdev/xray-core/v26/common"
	"github.com/xmplusdev/xray-core/v26/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
