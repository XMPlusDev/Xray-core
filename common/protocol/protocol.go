package protocol // import "github.com/xmplusdev/xray-core/v26/common/protocol"

import (
	"errors"
)

var ErrProtoNeedMoreData = errors.New("protocol matches, but need more data to complete sniffing")
