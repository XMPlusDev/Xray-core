package splithttp_test

import (
	"testing"

	. "github.com/xmplusdev/xray-core/v24/transport/internet/splithttp"
)

func Test_GetNormalizedPath(t *testing.T) {
	c := Config{
		Path: "/?world",
	}

	path := c.GetNormalizedPath()
	if path != "/" {
		t.Error("Unexpected: ", path)
	}
}
