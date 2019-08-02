package controllers

import (
	"TeamProject/pkg/models"
	"log"
)

func CreatePet(pet *models.Pet) {
		_, err := db.Exec("insert into pets (nickname, type,breed, age, weight, gender, description) values ($1,$2,$3,$4,$5,$6,$7)",
			pet.Nickname,pet.Type,pet.Breed,pet.Age,pet.Weight,pet.Gender,pet.Description)
		if err != nil {
			log.Println(err)
		}
	}

func DeletePet (pet *models.Pet){
	_, err := db.Exec("delete from pets where id = ?", pet.Id)
	if err != nil {
		log.Println(err)
	}
}

func GetPet (pet *models.Pet){
	_, err := db.Exec("select * from pets where id = ?", pet.Id)
	if err != nil {
		log.Println(err)
	}
}