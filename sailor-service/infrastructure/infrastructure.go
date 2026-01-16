package infrastructure

import (
	"github.com/google/wire"
	"github.com/kweaver-ai/dip-for-data-resource/sailor-service/infrastructure/repository/db"
)

var Set = wire.NewSet(
	db.NewData,
)
