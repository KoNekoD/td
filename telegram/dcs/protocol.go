package dcs

import (
	"net"

	"github.com/KoNekoD/td/transport"
)

type protocol interface {
	Handshake(conn net.Conn) (transport.Conn, error)
}
