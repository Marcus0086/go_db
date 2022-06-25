package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while intializing database", err)
	}
	allUsers := []defjson{}
	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error while reading users collection", err)
	}

	for _, user := range records {
		userfound := defjson{}
		if err := json.Unmarshal([]byte(user), &userfound); err != nil {
			fmt.Println("Error while unmashalling data", err)
		}
		allUsers = append(allUsers, userfound)
	}
	c.IndentedJSON(http.StatusOK, allUsers)

}

func getUser(c *gin.Context) {
	id := c.Param("id")
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while intializing database", err)
	}
	user := User{}
	record, err := db.Read("users", id, user)
	if err != nil {
		fmt.Println("Error while reading user", err)
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, record)
}

func createUser(c *gin.Context) {
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while intializing database", err)
	}
	newuser := User{}
	if err := c.BindJSON(&newuser); err != nil {
		println("Error while binding", err)
		return
	}
	db.Write("users", newuser.Name, User{
		Name:    newuser.Name,
		Age:     newuser.Age,
		Contact: newuser.Contact,
		Company: newuser.Company,
		Address: newuser.Address,
	})
	c.IndentedJSON(http.StatusCreated, newuser)
}

func deleteUser(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id not passed"})
		return
	}
	db, err := New("./", nil)
	if err != nil {
		fmt.Println("Error while intializing database", err)
	}
	if err := db.Delete("users", id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error happened while deleting user",
			"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted user :" + id})
}
