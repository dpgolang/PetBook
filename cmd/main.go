package main

import (
	"TeamProject/migrations"
	"TeamProject/pkg/driver"
	_ "TeamProject/pkg/handlers"
	handler "TeamProject/pkg/handlers"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/rubenv/sql-migrate"
	"net/http"
	"os"
)

//Init in folder - driver



func main() {
	db:= driver.ConnectDB()
	defer db.Close()

	mig := migrations.GetMigrations()
	migrations.MigrateUp(db,mig)


	router := mux.NewRouter()
	router.HandleFunc("/",handler.LoginHandler)
	router.HandleFunc("/registration",handler.CreateUserHandler)
	router.HandleFunc("/petcabinet",handler.CreatePetHandler)
	//router.HandleFunc("/vetcabinet",CreateVetHandler)
	router.HandleFunc("/main",handler.MainPageHandler)


	http.Handle("/", router)

	fmt.Println("Server is listening...")
	_ = http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, router))
}

