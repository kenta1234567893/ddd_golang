package libs

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kenta1234567893/upsider-coding-test/ent"
)

func GetClient() *ent.Client {
	// TODO 環境変数に切り出し
	client, err := ent.Open("mysql", "localuser:localpass@tcp(localhost:13306)/company?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return client
}
