//go:build windows
// +build windows

package client

import (
	"context"
	"net"
	"net/url"

	"github.com/Microsoft/go-winio"
)

func dialUnix(ctx context.Context, u *url.URL, dialer net.Dialer) (net.Conn, error) {
	return winio.DialPipeContext(ctx, u.Opaque)
}
