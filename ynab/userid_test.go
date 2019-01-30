package ynab

import (
	"testing"
    "fmt"
    "os"
    "encoding/json"
)

func TestUserID(t *testing.T) {
  blob := `"` + os.Getenv("YNAB_USER") + `"`
  var userid UserID
  if err := json.Unmarshal([]byte(blob), &userid); err != nil {
      t.Fatal(err)
  }
  fmt.Println(userid)
}
