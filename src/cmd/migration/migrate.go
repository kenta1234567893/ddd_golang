package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kenta1234567893/upsider-coding-test/ent/migrate"
	"github.com/kenta1234567893/upsider-coding-test/src/cmd/server/libs"
)

func main() {
	migration()
}

func migration() {
	client := libs.GetClient()
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
