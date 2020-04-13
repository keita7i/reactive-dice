package main

func main() {
	di := DI{}
	panic(di.Router().Run(":80"))
}
