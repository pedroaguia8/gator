package cli

import (
	"github.com/pedroaguia8/gator/internal/config"
	"github.com/pedroaguia8/gator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *config.Config
}
