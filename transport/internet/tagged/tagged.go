package tagged

import (
	"context"

	"github.com/xmplusdev/xray-core/v24/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
