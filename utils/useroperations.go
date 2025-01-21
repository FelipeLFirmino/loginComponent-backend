package utils

import (
	"encoding/json"
	"errors"

	"desafio-backend/models"
	"os"
)

// this function loads all the users from the file that sims the database
func LoadUsers(filename string) ([]models.User, error) {

	//open the file
	file, err := os.Open(filename)
	//if there is an error show it
	if err != nil {
		return nil, errors.New("an error occured while trying to open the file")
	}

	// create a var that is a slice of USERS(a struct) than use the decoder on it
	var users []models.User
	decoder := json.NewDecoder(file)
	//the error must be nil here, this means that the decoder was able to decode the file
	err = decoder.Decode(&users)
	//if there is an error at the decoding show it
	if err != nil {
		return nil, errors.New("an error occured while trying to decode")
	}
	return users, err

}

// this function saves an user to the file, it uses a byte slice of type user and the file name
func SaveUser(filename string, users []models.User) error {

	//it idents the byteslice of users(which in this case is a json) into a byteslice of bytes
	//so qe can use the writefile function that only accepts a slice of bytes
	data, err := json.MarshalIndent(users, "", " ")
	//if the error is different of nil log it
	if err != nil {
		return errors.New("an error occured while trying to indent ")
	}

	//Uses the filename and the data slice of bytes on the writefile function
	err = os.WriteFile(filename, data, 0666)
	//if the function returns an error different than new log it
	if err != nil {
		return errors.New("an error occured while trying to write in the file ")
	}

	return err
}

// this function calls the other two functions in it, it loads then checks if the user
// has tried to input an email or username that is already being used
func CheckAndSaveUser(filename string, newUser models.User) error {

	// calls the function load users
	users, err := LoadUsers(filename)
	//if there is an error log it
	if err != nil {
		return errors.New("an error occured while trying to load the users")
	}

	//iterate throught the users from the file and compare the email and username of each one with the new user
	for _, user := range users {
		if user.Email == newUser.Email {
			return errors.New("the email is already being used")
		} else if user.Username == newUser.Username {
			return errors.New("username is already being used")
		}
	}

	//adds the new user to the users (notice that they are byteslices of the user model)
	users = append(users, newUser)

	//call the save user function
	return SaveUser(filename, users)
}
