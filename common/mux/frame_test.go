package mux_test

import (
	"testing"

	"github.com/xmplusdev/xray-core/v26/common"
	"github.com/xmplusdev/xray-core/v26/common/buf"
	"github.com/xmplusdev/xray-core/v26/common/mux"
	"github.com/xmplusdev/xray-core/v26/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.example.com"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
