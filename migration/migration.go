package migration

import (
	"auth-service/db"
	"fmt"
)

func Migration() {
	database := db.Database()

	err := database.AutoMigrate()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("se migro")
	}
}
