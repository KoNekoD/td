package manager

import (
	"github.com/KoNekoD/td/bin"
	"github.com/KoNekoD/td/mtproto"
	"github.com/KoNekoD/td/tg"
)

// Handler abstracts updates and session handler.
type Handler interface {
	OnSession(cfg tg.Config, s mtproto.Session) error
	OnMessage(b *bin.Buffer) error
}

// NoopHandler is a noop handler.
type NoopHandler struct{}

// OnSession implements Handler.
func (n NoopHandler) OnSession(cfg tg.Config, s mtproto.Session) error {
	return nil
}

// OnMessage implements Handler
func (n NoopHandler) OnMessage(b *bin.Buffer) error {
	return nil
}
