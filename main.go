package main

//go:generate go-bindata -prefix "" -pkg migrations -ignore generated.go -o db/migrations/generated.go db/migrations

import (
	"database/sql"
	"fmt"

	"github.com/donkeysharp/events-manager/db/migrations"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	fmt.Println("Online Events Manager")
	fmt.Println(migrations.AssetNames())
	migration := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "db/migrations",
	}
	var connectionString string
	// TODO: Create configuration loader
	connectionString = "postgres://test:12345@localhost:5432/events_db?sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	result, err := migrate.Exec(db, "postgres", migration, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Migration result %d\n", result)
}
