//go:build !android

package boxdns

import (
	"context"
	"errors"
	D "github.com/sagernet/sing-dns"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	"net"
	"reflect"
	"runtime"
	"unsafe"
)

var underlyingDNS string

func init() {
	D.RegisterTransport([]string{"underlying"}, createUnderlyingTransport)
}

func createUnderlyingTransport(options D.TransportOptions) (D.Transport, error) {
	if runtime.GOOS != "windows" {
		// Linux no resolv.conf change
		return D.NewLocalTransport(D.TransportOptions{
			Context: options.Context,
			Logger:  options.Logger,
			Name:    options.Name,
			Dialer:  options.Dialer,
			Address: "local",
		}), nil
	}
	// Windows Underlying DNS hook
	t, _ := D.NewUDPTransport(options)
	handler_ := reflect.Indirect(reflect.ValueOf(t)).FieldByName("handler")
	handler_ = reflect.NewAt(handler_.Type(), unsafe.Pointer(handler_.UnsafeAddr())).Elem()
	handler_.Set(reflect.ValueOf(&myTransportHandler{t, options.Dialer}))
	return t, nil
}

var _ D.Transport = (*myTransportHandler)(nil)

type myTransportHandler struct {
	*D.UDPTransport
	dialer N.Dialer
}

func (t *myTransportHandler) DialContext(ctx context.Context) (net.Conn, error) {
	if underlyingDNS == "" {
		return nil, errors.New("no underlyingDNS")
	}
	return t.dialer.DialContext(ctx, "udp", M.ParseSocksaddrHostPort(underlyingDNS, 53))
}
