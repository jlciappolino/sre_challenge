package main

import (
	"github.com/jlciappolino/sre_challenge/apitools"
)

func main() {

	r := apitools.NewChallengeRouter()

	//handler := newHandler TODO

	//r.GET("/sold_items/:id", handler.Get)

	r.Run()
}
