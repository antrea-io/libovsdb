//go:build !windows
// +build !windows

package client

import (
	"context"
	"net"
	"net/url"
)

func dialUnix(ctx context.Context, u *url.URL, dialer net.Dialer) (net.Conn, error) {
	return dialer.DialContext(ctx, u.Scheme, u.Path)
}
