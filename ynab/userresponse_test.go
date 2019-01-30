package ynab

import (
	"testing"
    "fmt"
    "encoding/json"
)

func TestUserReponse(t *testing.T) {
    blob := `{"data": {"user": {"id": "` + os.Getenv("YNAB_USER") + `"}}}`
    var userResponse UserResponse
    if err := json.Unmarshal([]byte(blob), &userResponse); err != nil {
        t.Fatal(err)
    }
    fmt.Println(userResponse)
}
