package main

//go:generate go-bindata -prefix "" -pkg migrations -ignore generated.go -o db/migrations/generated.go db/migrations

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/donkeysharp/events-manager/db/migrations"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func testMigration(connectionString string, direction migrate.MigrationDirection) {
	fmt.Println(migrations.AssetNames())
	migration := &migrate.AssetMigrationSource{
		Asset:    migrations.Asset,
		AssetDir: migrations.AssetDir,
		Dir:      "db/migrations",
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	result, err := migrate.Exec(db, "postgres", migration, direction)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Migration result %d\n", result)
}

type event struct {
	gorm.Model
	Name        string
	Slug        string
	Description string
}

func testGorm(connectionString string) {
	fmt.Printf("Test Gorm with connection string %s\n", connectionString)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating an event...")
	db.Create(&event{Name: "event1", Slug: "slug_event_1", Description: "description"})

	fmt.Println("Retrieving all events...")
	var events []event
	result := db.Find(&events)
	fmt.Println(result.RowsAffected)
	fmt.Println(events[0].Slug)
}

func main() {
	var directionMap = make(map[string]migrate.MigrationDirection)
	directionMap["up"] = migrate.Up
	directionMap["down"] = migrate.Down

	fmt.Println("Online Events Manager")
	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	direction := migrateCmd.String("direction", "up", "Direction of the migration up/down")

	startCmd := flag.NewFlagSet("start", flag.ExitOnError)
	command := os.Args[1]

	var connectionString string
	connectionString = "postgres://test:12345@localhost:5432/events_db?sslmode=disable"

	switch command {
	case "migrate":
		migrateCmd.Parse(os.Args[2:])
		fmt.Printf("Migrating database. Direction: %s\n", *direction)
		testMigration(connectionString, directionMap[strings.ToLower(*direction)])
	case "start":
		startCmd.Parse(os.Args[2:])
		testGorm(connectionString)
	}
}
