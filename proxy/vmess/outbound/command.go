package outbound

import (

	"github.com/xmplusdev/xray-core/v26/common/net"
	"github.com/xmplusdev/xray-core/v26/common/protocol"
)

// As a stub command consumer.
func (h *Handler) handleCommand(dest net.Destination, cmd protocol.ResponseCommand) {
	switch cmd.(type) {
	default:
	}
}
