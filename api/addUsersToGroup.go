package api

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSV(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var dataArray []string
	for _, line := range lines {
		dataArray = append(dataArray, line[0])
	}

	return dataArray
}

func (c *OktaClient) GetUserIDs(arr []string) []string {
	arrayOfEmails := arr
	var arrayOfIds []string
	for _, email := range arrayOfEmails {
		user, _, err := c.User.GetUser(email, nil)
		if err != nil {
			log.Fatal(err)
		}
		arrayOfIds = append(arrayOfIds, user.Id)
	}
	return arrayOfIds
}

func (c *OktaClient) AddIDsToGroup(arr []string, groupID string) {
	var errors []error
	for _, userID := range arr {
		_, err := c.Group.AddUserToGroup(groupID, userID, nil)
		if err != nil {
			errors = append(errors, err)
		}
	}
	if errors != nil {
		log.Fatal(errors)
	}
}
