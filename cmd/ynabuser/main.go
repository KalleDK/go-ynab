package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kalledk/go-ynab/ynab"
	"github.com/kalledk/go-ynab/ynab/client"
	"github.com/kalledk/go-ynab/ynab/user"
)

func main() {
	accessToken, err := ynab.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		log.Fatalf("invalid token, %v", err)
	}
	c := client.NewClient(accessToken)

	myuser, err := user.GetUser(c)
	if err != nil {
		log.Fatalf("invalid response %v", err)
	}

	json, err := json.MarshalIndent(myuser, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}
