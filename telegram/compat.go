package telegram

import (
	"github.com/KoNekoD/td/session"
	"github.com/KoNekoD/td/tgerr"
)

// SessionStorage is alias of mtproto.SessionStorage.
type SessionStorage = session.Storage

// FileSessionStorage is alias of mtproto.FileSessionStorage.
type FileSessionStorage = session.FileStorage

// Error represents RPC error returned to request.
type Error = tgerr.Error
