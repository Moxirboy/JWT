package conn

import (
	"JWT/internal/configs"
)

var Db, _ = configs.NewPostgresConfig(configs.Configuration())
