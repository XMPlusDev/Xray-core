package taggedimpl

import (
	"context"

	"github.com/xmplusdev/xray-core/v24/common/errors"
	"github.com/xmplusdev/xray-core/v24/common/net"
	"github.com/xmplusdev/xray-core/v24/common/net/cnc"
	"github.com/xmplusdev/xray-core/v24/common/session"
	"github.com/xmplusdev/xray-core/v24/core"
	"github.com/xmplusdev/xray-core/v24/features/routing"
	"github.com/xmplusdev/xray-core/v24/transport/internet/tagged"
)

func DialTaggedOutbound(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error) {
	if core.FromContext(ctx) == nil {
		return nil, errors.New("Instance context variable is not in context, dial denied. ")
	}
	content := new(session.Content)
	content.SkipDNSResolve = true

	ctx = session.ContextWithContent(ctx, content)
	ctx = session.SetForcedOutboundTagToContext(ctx, tag)

	r, err := dispatcher.Dispatch(ctx, dest)
	if err != nil {
		return nil, err
	}
	var readerOpt cnc.ConnectionOption
	if dest.Network == net.Network_TCP {
		readerOpt = cnc.ConnectionOutputMulti(r.Reader)
	} else {
		readerOpt = cnc.ConnectionOutputMultiUDP(r.Reader)
	}
	return cnc.NewConnection(cnc.ConnectionInputMulti(r.Writer), readerOpt), nil
}

func init() {
	tagged.Dialer = DialTaggedOutbound
}
