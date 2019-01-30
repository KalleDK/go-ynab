package ynab

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestUserID(t *testing.T) {
	blob := `"` + os.Getenv("YNAB_USER") + `"`
	var userid UserID
	if err := json.Unmarshal([]byte(blob), &userid); err != nil {
		t.Fatal(err)
	}
	fmt.Println(userid)
}
