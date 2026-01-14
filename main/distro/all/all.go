package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/xmplusdev/xray-core/v26/app/dispatcher"
	_ "github.com/xmplusdev/xray-core/v26/app/proxyman/inbound"
	_ "github.com/xmplusdev/xray-core/v26/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/xmplusdev/xray-core/v26/app/commander"
	_ "github.com/xmplusdev/xray-core/v26/app/log/command"
	_ "github.com/xmplusdev/xray-core/v26/app/proxyman/command"
	_ "github.com/xmplusdev/xray-core/v26/app/stats/command"

	// Developer preview services
	_ "github.com/xmplusdev/xray-core/v26/app/observatory/command"

	// Other optional features.
	_ "github.com/xmplusdev/xray-core/v26/app/dns"
	_ "github.com/xmplusdev/xray-core/v26/app/dns/fakedns"
	_ "github.com/xmplusdev/xray-core/v26/app/log"
	_ "github.com/xmplusdev/xray-core/v26/app/metrics"
	_ "github.com/xmplusdev/xray-core/v26/app/policy"
	_ "github.com/xmplusdev/xray-core/v26/app/reverse"
	_ "github.com/xmplusdev/xray-core/v26/app/router"
	_ "github.com/xmplusdev/xray-core/v26/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/xmplusdev/xray-core/v26/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/xmplusdev/xray-core/v26/proxy/blackhole"
	_ "github.com/xmplusdev/xray-core/v26/proxy/dns"
	_ "github.com/xmplusdev/xray-core/v26/proxy/dokodemo"
	_ "github.com/xmplusdev/xray-core/v26/proxy/freedom"
	_ "github.com/xmplusdev/xray-core/v26/proxy/http"
	_ "github.com/xmplusdev/xray-core/v26/proxy/loopback"
	_ "github.com/xmplusdev/xray-core/v26/proxy/shadowsocks"
	_ "github.com/xmplusdev/xray-core/v26/proxy/socks"
	_ "github.com/xmplusdev/xray-core/v26/proxy/trojan"
	_ "github.com/xmplusdev/xray-core/v26/proxy/vless/inbound"
	_ "github.com/xmplusdev/xray-core/v26/proxy/vless/outbound"
	_ "github.com/xmplusdev/xray-core/v26/proxy/vmess/inbound"
	_ "github.com/xmplusdev/xray-core/v26/proxy/vmess/outbound"
	_ "github.com/xmplusdev/xray-core/v26/proxy/wireguard"

	// Transports
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/grpc"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/httpupgrade"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/kcp"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/reality"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/splithttp"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/tcp"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/tls"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/udp"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/websocket"

	// Transport headers
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/headers/http"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/headers/noop"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/headers/srtp"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/headers/tls"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/headers/utp"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/headers/wechat"
	_ "github.com/xmplusdev/xray-core/v26/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/xmplusdev/xray-core/v26/main/json"
	_ "github.com/xmplusdev/xray-core/v26/main/toml"
	_ "github.com/xmplusdev/xray-core/v26/main/yaml"

	// Load config from file or http(s)
	_ "github.com/xmplusdev/xray-core/v26/main/confloader/external"

	// Commands
	_ "github.com/xmplusdev/xray-core/v26/main/commands/all"
)
