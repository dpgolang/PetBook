package handler

import (
	"net/http"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request){
	//db:=driver.ConnectDB()
	////pets:= []models.Pet{}
	//row, err := db.Query("select * from pets where id = ?", 3)
	//if err != nil {
	//	log.Println(err)
	//}
	//pet:=models.Pet{}
	//err = row.Scan(&pet.Id, &pet.Nickname, &pet.Type, &pet.Breed, &pet.Age,&pet.Weight,&pet.Gender,&pet.Description)
	//if err != nil {
	//	log.Println(err)
	//	http.Error(w, http.StatusText(404), http.StatusNotFound)
	//}
	//defer row.Close()
	//
	//tmpl, _ := template.ParseFiles("web/info_Pet.html")
	//tmpl.Execute(w, pet)

	http.ServeFile(w, r, "web/main.html")

}
