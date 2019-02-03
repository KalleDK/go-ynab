package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kalledk/go-ynab/ynab/api"
	"github.com/kalledk/go-ynab/ynab/client"
)

func main() {
	accessToken, err := api.NewAccessToken(os.Getenv("YNAB_TOKEN"))
	if err != nil {
		log.Fatalf("invalid token, %v", err)
	}
	c := client.NewClient(accessToken)

	myuser, err := c.User().Get()
	if err != nil {
		log.Fatalf("invalid response %v", err)
	}

	json, err := json.MarshalIndent(myuser, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}
