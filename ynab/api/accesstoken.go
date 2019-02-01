package api

import (
	"encoding/hex"
	"fmt"
)

type AccessToken [32]byte

func (a AccessToken) String() string {
	var buf [64]byte
	hex.Encode(buf[:], a[:])
	return string(buf[:])
}

func ParseAccessToken(b []byte) (a AccessToken, err error) {
	n, err := hex.Decode(a[:], b)
	if n != 32 {
		return a, fmt.Errorf("invalid token length %v", n)
	}
	return a, nil
}

func NewAccessToken(s string) (accessToken AccessToken, err error) {
	accessToken, err = ParseAccessToken([]byte(s))
	if err != nil {
		return
	}

	return
}
