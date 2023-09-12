package data

import (
	"go-grpc/internal/data/article"

	"github.com/google/wire"
)

var Provider = wire.NewSet(article.Provider)
