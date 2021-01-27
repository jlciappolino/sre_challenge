package main

func main() {

	r := apitools.NewChallengeRouter()

	//handler := newHandler TODO

	//r.GET("/items/:id", handler.Get)

	r.Run()
}
