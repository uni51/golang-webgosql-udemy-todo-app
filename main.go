package main

import (
	"fmt"

	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

	// user, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(user)

	// sessioon, err := user.CreateSession()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(sessioon)

	// valid, _ := sessioon.CheckSession()
	// fmt.Println(valid)
}
