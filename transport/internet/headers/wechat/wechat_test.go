package wechat_test

import (
	"context"
	"testing"

	"github.com/xmplusdev/xray-core/v25/common"
	"github.com/xmplusdev/xray-core/v25/common/buf"
	. "github.com/xmplusdev/xray-core/v25/transport/internet/headers/wechat"
)

func TestUTPWrite(t *testing.T) {
	videoRaw, err := NewVideoChat(context.Background(), &VideoConfig{})
	common.Must(err)

	video := videoRaw.(*VideoChat)

	payload := buf.New()
	video.Serialize(payload.Extend(video.Size()))

	if payload.Len() != video.Size() {
		t.Error("expected payload size ", video.Size(), " but got ", payload.Len())
	}
}
