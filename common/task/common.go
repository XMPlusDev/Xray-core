package task

import "github.com/xmplusdev/xray-core/v24/common"

// Close returns a func() that closes v.
func Close(v interface{}) func() error {
	return func() error {
		return common.Close(v)
	}
}
