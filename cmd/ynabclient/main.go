package main

import (
	"github.com/kalledk/go-ynab/ynab"
)

func main() {
	c := ynab.Client{"Flaf"}
	c.Demo()
}
