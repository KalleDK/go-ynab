package ynab

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	blob := `{"id": "` + os.Getenv("YNAB_USER") + `"}`
	var user User
	if err := json.Unmarshal([]byte(blob), &user); err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}
