package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/KalleDK/go-ynab/pkg/ynab"
)

type teeReader struct {
	io.Closer
	io.Reader
}

func (t teeReader) Close() error {
	os.Stderr.Write([]byte{'\n'})
	os.Stderr.Sync()
	return t.Closer.Close()
}

type wrapped struct {
	base http.RoundTripper
}

func (w wrapped) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Print(r.URL)
	if r.Body != nil {
		r.Body = teeReader{r.Body, io.TeeReader(r.Body, os.Stderr)}
	}
	resp, err := w.base.RoundTrip(r)
	if err != nil {
		return resp, err
	}
	fmt.Fprintf(os.Stderr, "StausCode: %v %s\n", resp.StatusCode, resp.Status)
	/*
		for k, v := range resp.Header {
			fmt.Fprintf(os.Stderr, "%s: %v\n", k, v)
		}
	*/

	resp.Body = teeReader{resp.Body, io.TeeReader(resp.Body, os.Stderr)}
	return resp, err
}

func main() {
	token := ynab.Token(os.Getenv("TOKEN"))

	client := ynab.Config{
		Client: &http.Client{
			Transport: wrapped{http.DefaultTransport},
		},
		Token: token,
	}.NewClient()

	/*
		budgetID := uuid.MustParse(os.Getenv("BUDGETID"))

		payees, err := client.GetPayees(budgetID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", payees)

		accounts, err := client.GetAccounts(budgetID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", accounts)

		catGroups, err := client.GetCategories(budgetID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", catGroups)
	*/
	user, err := client.GetUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User %+v\n", user)

	budgets, err := client.GetBudgets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Budgets %+v\n", budgets)

}
