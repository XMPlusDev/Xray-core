package dns

import (
	"context"

	"github.com/xmplusdev/xray-core/common/errors"
	"github.com/xmplusdev/xray-core/common/net"
	"github.com/xmplusdev/xray-core/core"
	"github.com/xmplusdev/xray-core/features/dns"
)

type FakeDNSServer struct {
	fakeDNSEngine dns.FakeDNSEngine
}

func NewFakeDNSServer() *FakeDNSServer {
	return &FakeDNSServer{}
}

func (FakeDNSServer) Name() string {
	return "FakeDNS"
}

func (f *FakeDNSServer) QueryIP(ctx context.Context, domain string, _ net.IP, opt dns.IPOption, _ bool) ([]net.IP, error) {
	if f.fakeDNSEngine == nil {
		if err := core.RequireFeatures(ctx, func(fd dns.FakeDNSEngine) {
			f.fakeDNSEngine = fd
		}); err != nil {
			return nil, errors.New("Unable to locate a fake DNS Engine").Base(err).AtError()
		}
	}
	var ips []net.Address
	if fkr0, ok := f.fakeDNSEngine.(dns.FakeDNSEngineRev0); ok {
		ips = fkr0.GetFakeIPForDomain3(domain, opt.IPv4Enable, opt.IPv6Enable)
	} else {
		ips = f.fakeDNSEngine.GetFakeIPForDomain(domain)
	}

	netIP, err := toNetIP(ips)
	if err != nil {
		return nil, errors.New("Unable to convert IP to net ip").Base(err).AtError()
	}

	errors.LogInfo(ctx, f.Name(), " got answer: ", domain, " -> ", ips)

	if len(netIP) > 0 {
		return netIP, nil
	}
	return nil, dns.ErrEmptyResponse
}
