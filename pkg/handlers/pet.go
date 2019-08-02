package handler

import (
	"TeamProject/pkg/controllers"
	"TeamProject/pkg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreatePetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		newp:=&models.Pet{}
		newp.Nickname = r.FormValue("nickname")
		newp.Type = r.FormValue("pet-type")
		newp.Breed = r.FormValue("breed")
		newp.Age = r.FormValue("age")
		newp.Weight = r.FormValue("weight")
		newp.Gender = r.FormValue("gender")
		newp.Description = r.FormValue("description")

		controllers.CreatePet(newp)


		http.Redirect(w, r, "/main", 301)
	} else {
		http.ServeFile(w, r, "web/cabinetPet.html")
	}
}

func DeletePetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pet:= &models.Pet{}
	pet.Id = vars["id"]
	controllers.DeletePet(pet)
	http.Redirect(w, r, "/login.html", 301)
}

//func GetPetHandler(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	pet:= &models.Pet{}
//	pet.Id = vars["id"]
//
//	controllers.GetPet(pet)
//	err := row.Scan(&pet.Id, &pet.Nickname, &pet.Type, &pet.Breed, &pet.Age,&pet.Weight,&pet.Gender,&pet.Description)
//	if err != nil {
//		log.Println(err)
//		http.Error(w, http.StatusText(404), http.StatusNotFound)
//	} else {
//		tmpl, _ := template.ParseFiles("web/edit_Pet.html")
//		tmpl.Execute(w, pet)
//	}
//}

//func UpdatePetHandler(w http.ResponseWriter, r *http.Request) {
//	err := r.ParseForm()
//	if err != nil {
//		log.Println(err)
//	}
//	id := r.FormValue("id")
//	nickname := r.FormValue("nickname")
//	petType := r.FormValue("pet-type")
//	breed := r.FormValue("breed")
//	age := r.FormValue("age")
//	weight:= r.FormValue("weight")
//	gender:= r.FormValue("gender")
//	petDescription:= r.FormValue("description")
//
//	_, err = db.Exec("update petbook.pets set nickname=?, type=?,breed=?, age = ?,weight=?,gender=?,description=? where id = ?",
//		nickname, petType, breed, age,weight,gender,petDescription, id)
//
//	if err != nil {
//		log.Println(err)
//	}
//	http.Redirect(w, r, "/", 301)
//}


