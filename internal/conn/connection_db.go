package conn

import (
	"JWT/internal/configs"
	"JWT/pkg/postgres"
)

var Db, _ = postgres.NewPostgresConfig(configs.Configuration())
