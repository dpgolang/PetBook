package controllers

import (
	"TeamProject/pkg/driver"
	"TeamProject/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var db=driver.ConnectDB()

func Login(u *models.User) bool{
	row := db.QueryRow("select email,password from users")
	user := models.User{}
	err := row.Scan(&user.Email, &user.Password)
	if err != nil {
		log.Println(err)
	}
	if user.Email == u.Email && user.Password == u.Password {
		return true
	}else {return false}

}

func Register (u *models.User){
	_,err:= db.Exec("INSERT INTO users (name, surname, email, role, password)values ($1,$2,$3,$4,$5)",
		u.FirstName,u.LastName,u.Email,u.Role,u.Password)
	if err != nil {
		log.Println(err)
	}
}



func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func GetUser(db *sqlx.DB, user *models.User) error {
	err := db.QueryRowx("select * from users where email=$1", user.Email).StructScan(user)
	if err != nil {
		logErr(err)
		return fmt.Errorf("cannot scan user from db: %v", err)
	}
	return nil
}

func GetUsers(db *sqlx.DB) ([]models.User, error) {
	rows, err := db.Query("select * from users")
	logErr(err)
	defer rows.Close()
	users := []models.User{}
	err = sqlx.StructScan(rows, &users)
	if err != nil {
		logErr(err)
		return users, fmt.Errorf("cannot scan users from db: %v", err)
	}
	return users, nil
}


func ChangePassword(db *sqlx.DB, user *models.User, newPassword string) error {
	res, err := db.Exec("UPDATE users SET password=$1 WHERE email = $2",
		newPassword, user.Email)
	if err != nil {
		return fmt.Errorf("cannot update users in db: %v", err)
	}
	num, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot affect rows in users in db: %v", err)
	}
	if num != 1 {
		return fmt.Errorf("cannot find this user")
	}
	return nil
}



