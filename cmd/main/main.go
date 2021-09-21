package main

import (
	"net/http"
	"swaggertest/plugins/api"
)

func main() {

	http.ListenAndServe(":9944", api.NewApi())
}
