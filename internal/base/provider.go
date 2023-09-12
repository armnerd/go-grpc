package base

import "github.com/google/wire"

var Provider = wire.NewSet(ConfProvider, LoggerProvider, DbProvider, CacheProvider)
