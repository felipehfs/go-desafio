package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/felipehfs/godesafio/controllers"
	"github.com/felipehfs/godesafio/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "godesafio"
)

func main() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", sqlInfo)

	if err != nil {
		panic(err)
	}

	jwt := utils.EnabledJwt()

	userHandler := controllers.UserHandler{DB: db}

	defer db.Close()
	mux := mux.NewRouter()
	mux.HandleFunc("/user", userHandler.Register).Methods("POST")
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ping"))
	}).Methods("GET")

	mux.HandleFunc("/user/{uuid}", jwt(userHandler.FindByUUID)).Methods("GET")
	mux.HandleFunc("/user", jwt(userHandler.RemoveUser)).Methods("DELETE")
	mux.HandleFunc("/user", jwt(userHandler.UpdateUser)).Methods("PUT")

	cors := utils.Cors()
	mux.HandleFunc("/auth", userHandler.SignIn).Methods("POST")
	http.ListenAndServe(":8080", cors(mux))
}
