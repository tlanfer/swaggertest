package api

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "swaggertest/plugins/api/docs"
	"swaggertest/plugins/api/index"
)

//go:generate swag init -g ./api.go

// @title Swagger Example API
// @version 1.0
// @license.name Apache 2.0

func NewApi() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", index.NewHandler())
	mux.Handle("/swagger/", httpSwagger.Handler())

	return mux
}
