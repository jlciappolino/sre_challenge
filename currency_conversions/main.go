package main

func main() {

	r := apitools.NewChallengeRouter()

	//handler := newHandler TODO

	//r.GET("/currency_conversions/:id", handler.Get)

	r.Run()
}
