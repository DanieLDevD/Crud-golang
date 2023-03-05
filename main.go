package main

import (
	"fmt"
	"net/http"

	"github.com/DanielDevD/Crud-golang/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func main(){
	viper.SetConfigFile("./envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}