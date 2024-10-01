package shared

import (
	"fmt"

	"github.com/kenta1234567893/upsider-coding-test/ent"
)

func GetClient() (*ent.Client, error) {
	client, err := ent.Open("mysql", "localuser:localpass@tcp(localhost:13306)/company?parseTime=True")
	if err != nil {
		return nil, fmt.Errorf("ERROR: failed opening connection to mysql: " + err.Error())
	}

	return client, nil
}
