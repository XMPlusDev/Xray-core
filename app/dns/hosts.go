package dns

import (
	"context"

	"github.com/xmplusdev/xray-core/v24/common/errors"
	"github.com/xmplusdev/xray-core/v24/common/net"
	"github.com/xmplusdev/xray-core/v24/common/strmatcher"
	"github.com/xmplusdev/xray-core/v24/features/dns"
)

// StaticHosts represents static domain-ip mapping in DNS server.
type StaticHosts struct {
	ips      [][]net.Address
	matchers *strmatcher.MatcherGroup
}

// NewStaticHosts creates a new StaticHosts instance.
func NewStaticHosts(hosts []*Config_HostMapping) (*StaticHosts, error) {
	g := new(strmatcher.MatcherGroup)
	sh := &StaticHosts{
		ips:      make([][]net.Address, len(hosts)+16),
		matchers: g,
	}

	for _, mapping := range hosts {
		matcher, err := toStrMatcher(mapping.Type, mapping.Domain)
		if err != nil {
			return nil, errors.New("failed to create domain matcher").Base(err)
		}
		id := g.Add(matcher)
		ips := make([]net.Address, 0, len(mapping.Ip)+1)
		switch {
		case len(mapping.ProxiedDomain) > 0:
			ips = append(ips, net.DomainAddress(mapping.ProxiedDomain))
		case len(mapping.Ip) > 0:
			for _, ip := range mapping.Ip {
				addr := net.IPAddress(ip)
				if addr == nil {
					return nil, errors.New("invalid IP address in static hosts: ", ip).AtWarning()
				}
				ips = append(ips, addr)
			}
		default:
			return nil, errors.New("neither IP address nor proxied domain specified for domain: ", mapping.Domain).AtWarning()
		}

		sh.ips[id] = ips
	}

	return sh, nil
}

func filterIP(ips []net.Address, option dns.IPOption) []net.Address {
	filtered := make([]net.Address, 0, len(ips))
	for _, ip := range ips {
		if (ip.Family().IsIPv4() && option.IPv4Enable) || (ip.Family().IsIPv6() && option.IPv6Enable) {
			filtered = append(filtered, ip)
		}
	}
	return filtered
}

func (h *StaticHosts) lookupInternal(domain string) []net.Address {
	var ips []net.Address
	for _, id := range h.matchers.Match(domain) {
		ips = append(ips, h.ips[id]...)
	}
	return ips
}

func (h *StaticHosts) lookup(domain string, option dns.IPOption, maxDepth int) []net.Address {
	switch addrs := h.lookupInternal(domain); {
	case len(addrs) == 0: // Not recorded in static hosts, return nil
		return nil
	case len(addrs) == 1 && addrs[0].Family().IsDomain(): // Try to unwrap domain
		errors.LogDebug(context.Background(), "found replaced domain: ", domain, " -> ", addrs[0].Domain(), ". Try to unwrap it")
		if maxDepth > 0 {
			unwrapped := h.lookup(addrs[0].Domain(), option, maxDepth-1)
			if unwrapped != nil {
				return unwrapped
			}
		}
		return addrs
	default: // IP record found, return a non-nil IP array
		return filterIP(addrs, option)
	}
}

// Lookup returns IP addresses or proxied domain for the given domain, if exists in this StaticHosts.
func (h *StaticHosts) Lookup(domain string, option dns.IPOption) []net.Address {
	return h.lookup(domain, option, 5)
}
