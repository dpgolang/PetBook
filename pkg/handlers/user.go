package handler

import (
	"TeamProject/pkg/controllers"
	"TeamProject/pkg/models"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		}

		logUser:=&models.User{}
		logUser.Email = r.FormValue("email")
		logUser.Password = r.FormValue("password")

		if controllers.Login(logUser) {
			http.Redirect(w, r, "/main", 303)
		} else {
			http.ServeFile(w, r, "web/login.html")
		}
	} else {
		http.ServeFile(w, r, "web/login.html")
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		regUser:=&models.User{}
		regUser.FirstName = r.FormValue("name")
		regUser.LastName = r.FormValue("surname")
		regUser.Email = r.FormValue("email")
		regUser.Role = r.FormValue("role")
		regUser.Password = r.FormValue("password")
		confirmation := r.FormValue("confirmation")

		if regUser.Password!=confirmation {
			http.ServeFile(w, r, "web/registration.html")
		}
		controllers.Register(regUser)
		switch regUser.Role {
		case "Pet":
			http.Redirect(w, r, "/petcabinet", 303)
		case "Veterinary":
			http.Redirect(w, r, "/vetcabinet", 303)
		}
	} else {
		http.ServeFile(w, r, "web/registration.html")
	}
}
