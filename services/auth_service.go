package services

import (
	"encoding/json"
	"os"
	"github.com/gizemcetincaglar/todo_app/models"
)

var Users []models.User

func LoadUsersFromFile() error {
	file, err := os.Open("mockdata/users.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&Users)
}

func FindUserByUsername(username string) *models.User {
	for _, user := range Users {
		if user.Username == username {
			return &user
		}
	}
	return nil
}
