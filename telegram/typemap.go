package telegram

import (
	"sync"

	"github.com/KoNekoD/td/mt"
	"github.com/KoNekoD/td/proto"
	"github.com/KoNekoD/td/tg"
	"github.com/KoNekoD/td/tmap"
)

// Port is default port used by telegram.
const Port = 443

var (
	typesMap  *tmap.Map
	typesOnce sync.Once
)

func getTypesMapping() *tmap.Map {
	typesOnce.Do(func() {
		typesMap = tmap.New(
			tg.TypesMap(),
			mt.TypesMap(),
			proto.TypesMap(),
		)
	})
	return typesMap
}
