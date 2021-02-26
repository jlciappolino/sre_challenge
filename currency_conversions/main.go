package main

import (
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/currency_conversions/internal"
)

func main() {

	r := apitools.NewChallengeRouter()

	handler := internal.NewCurrencyController()

	r.GET("/currency_conversions/:id", handler.Get)

	r.Run()
}
