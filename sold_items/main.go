package main

func main() {

	r := apitools.NewChallengeRouter()

	//handler := newHandler TODO

	//r.GET("/sold_items/:id", handler.Get)

	r.Run()
}
