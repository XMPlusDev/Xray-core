package all

import (
	"github.com/xmplusdev/xray-core/v24/main/commands/all/api"
	"github.com/xmplusdev/xray-core/v24/main/commands/all/convert"
	"github.com/xmplusdev/xray-core/v24/main/commands/all/tls"
	"github.com/xmplusdev/xray-core/v24/main/commands/base"
)

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		convert.CmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
