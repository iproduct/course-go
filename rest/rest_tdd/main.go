package main

import "github.com/iproduct/coursego/rest/rest_tdd/rest"

func main() {
	a := rest.App{}
	a.Initialize("root", "root", "go_rest_api")

	a.Run(":8080")
}
