package builder

import (
	database "github.com/MonikaCat/ibcjuno/db"
	"github.com/MonikaCat/ibcjuno/db/postgresql"
)

// DatabaseBuilder represents ConnectDatabase implementation that builds database
// instance based on the configuration set inside config.yaml file
func DatabaseBuilder(ctx *database.DatabaseContext) (database.Database, error) {
	return postgresql.ConnectDatabase(ctx)
}
