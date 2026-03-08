package main

import (
	"github.com/go-minstack/core"
	"github.com/go-minstack/migration"
	"github.com/go-minstack/sqlite"

	"example/migration-hello/migrations"
)

func main() {
	app := core.New(
		sqlite.Module(),
		migration.Module(migrations.FS),
	)
	app.Invoke(migration.Run)
	app.Run()
}
