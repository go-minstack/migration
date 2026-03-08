package main

import (
	"embed"

	"github.com/go-minstack/core"
	"github.com/go-minstack/migration"
	"github.com/go-minstack/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func main() {
	app := core.New(
		sqlite.Module(),
		migration.Module(migrationsFS),
	)
	app.Invoke(migration.Run)
	app.Run()
}
