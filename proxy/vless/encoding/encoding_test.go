package encoding_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xmplusdev/xray-core/v25/common"
	"github.com/xmplusdev/xray-core/v25/common/buf"
	"github.com/xmplusdev/xray-core/v25/common/net"
	"github.com/xmplusdev/xray-core/v25/common/protocol"
	"github.com/xmplusdev/xray-core/v25/common/uuid"
	"github.com/xmplusdev/xray-core/v25/proxy/vless"
	. "github.com/xmplusdev/xray-core/v25/proxy/vless/encoding"
)

func toAccount(a *vless.Account) protocol.Account {
	account, err := a.AsAccount()
	common.Must(err)
	return account
}

func TestRequestSerialization(t *testing.T) {
	user := &protocol.MemoryUser{
		Level: 0,
		Email: "test@example.com",
	}
	id := uuid.New()
	account := &vless.Account{
		Id: id.String(),
	}
	user.Account = toAccount(account)

	expectedRequest := &protocol.RequestHeader{
		Version: Version,
		User:    user,
		Command: protocol.RequestCommandTCP,
		Address: net.DomainAddress("www.example.com"),
		Port:    net.Port(443),
	}
	expectedAddons := &Addons{}

	buffer := buf.StackNew()
	common.Must(EncodeRequestHeader(&buffer, expectedRequest, expectedAddons))

	Validator := new(vless.MemoryValidator)
	Validator.Add(user)

	actualRequest, actualAddons, _, err := DecodeRequestHeader(false, nil, &buffer, Validator)
	common.Must(err)

	if r := cmp.Diff(actualRequest, expectedRequest, cmp.AllowUnexported(protocol.ID{})); r != "" {
		t.Error(r)
	}

	addonsComparer := func(x, y *Addons) bool {
		return (x.Flow == y.Flow) && (cmp.Equal(x.Seed, y.Seed))
	}
	if r := cmp.Diff(actualAddons, expectedAddons, cmp.Comparer(addonsComparer)); r != "" {
		t.Error(r)
	}
}

func TestInvalidRequest(t *testing.T) {
	user := &protocol.MemoryUser{
		Level: 0,
		Email: "test@example.com",
	}
	id := uuid.New()
	account := &vless.Account{
		Id: id.String(),
	}
	user.Account = toAccount(account)

	expectedRequest := &protocol.RequestHeader{
		Version: Version,
		User:    user,
		Command: protocol.RequestCommand(100),
		Address: net.DomainAddress("www.example.com"),
		Port:    net.Port(443),
	}
	expectedAddons := &Addons{}

	buffer := buf.StackNew()
	common.Must(EncodeRequestHeader(&buffer, expectedRequest, expectedAddons))

	Validator := new(vless.MemoryValidator)
	Validator.Add(user)

	_, _, _, err := DecodeRequestHeader(false, nil, &buffer, Validator)
	if err == nil {
		t.Error("nil error")
	}
}

func TestMuxRequest(t *testing.T) {
	user := &protocol.MemoryUser{
		Level: 0,
		Email: "test@example.com",
	}
	id := uuid.New()
	account := &vless.Account{
		Id: id.String(),
	}
	user.Account = toAccount(account)

	expectedRequest := &protocol.RequestHeader{
		Version: Version,
		User:    user,
		Command: protocol.RequestCommandMux,
		Address: net.DomainAddress("v1.mux.cool"),
	}
	expectedAddons := &Addons{}

	buffer := buf.StackNew()
	common.Must(EncodeRequestHeader(&buffer, expectedRequest, expectedAddons))

	Validator := new(vless.MemoryValidator)
	Validator.Add(user)

	actualRequest, actualAddons, _, err := DecodeRequestHeader(false, nil, &buffer, Validator)
	common.Must(err)

	if r := cmp.Diff(actualRequest, expectedRequest, cmp.AllowUnexported(protocol.ID{})); r != "" {
		t.Error(r)
	}

	addonsComparer := func(x, y *Addons) bool {
		return (x.Flow == y.Flow) && (cmp.Equal(x.Seed, y.Seed))
	}
	if r := cmp.Diff(actualAddons, expectedAddons, cmp.Comparer(addonsComparer)); r != "" {
		t.Error(r)
	}
}
