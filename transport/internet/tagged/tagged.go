package tagged

import (
	"context"

	"github.com/xmplusdev/xray-core/v25/common/net"
	"github.com/xmplusdev/xray-core/v25/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
