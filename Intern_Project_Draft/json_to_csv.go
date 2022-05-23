package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("trial_users.json") //jsonFile holds what is in our opened json file while err stores an error message if opening of the file failed.

	//We need to handle any error in opening the file
	if err != nil {
		fmt.Println(err)
	}
	//if succesfully opened,
	fmt.Println("Successfully opened jsonfile.")

	//Defer closing opened jsonfile till we are done parsing
	defer jsonFile.Close()

	//read our opened jsonFile as a byte array that we use the json unmarshall method on.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users //Initialize our users Struct defined above

	// we unmarshal our byteArray which contains our jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	//Creating a csv file to put the parsed data into
	//Column data that will be entered into csv
	userData := [][]string{}
	row1 := []string{"Name", "Age", "Type", "Facebook Url", "Twitter Url"}
	userData = append(userData, row1)
	for i := 0; i < len(users.Users); i++ {
		row2 := []string{users.Users[i].Name, strconv.Itoa(users.Users[i].Age), users.Users[i].Type, users.Users[i].Social.Facebook, users.Users[i].Social.Twitter}
		userData = append(userData, row2)
	}
	//fmt.Println(userData)

	csvFile, err := os.Create("user.csv") //attempting to create csv
	//erro handling in creating csv
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile) //Writing into created "user.csv"

	//filling up csv file
	for _, userRow := range userData {
		_ = csvwriter.Write(userRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
