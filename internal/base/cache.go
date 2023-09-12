package base

import "github.com/google/wire"

var CacheProvider = wire.NewSet(NewCache)

type Cache struct {
}

func NewCache() Cache {
	return Cache{}
}
