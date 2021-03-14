package mtproto

import (
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/proto"
)

func (c *Conn) handleFutureSalts(b *bin.Buffer) error {
	var res proto.FutureSalts

	if err := res.Decode(b); err != nil {
		return xerrors.Errorf("error decode: %w", err)
	}

	c.saltsMux.Lock()
	c.salts = append(c.salts, res.Salts...)
	c.saltsMux.Unlock()

	return nil
}
