package udp

import (
	"github.com/xmplusdev/xray-core/v24/common/buf"
	"github.com/xmplusdev/xray-core/v24/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
