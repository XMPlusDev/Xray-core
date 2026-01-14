package tagged

import (
	"context"

	"github.com/xmplusdev/xray-core/v26/common/net"
	"github.com/xmplusdev/xray-core/v26/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
