package ynab

import (
	"testing"
)

func TestAccountID(t *testing.T) {
  var access = os.Getenv("YNAB_ACCOUNT")
  a, _ := NewAccountID(access)
  if access != a.String() {
    t.Errorf("Expected %v got %v", access, a.String())
  }
}
