package mtproto

import "github.com/KoNekoD/td/proto"

func (c *Conn) newMessageID() int64 {
	return c.messageID.New(proto.MessageFromClient)
}
