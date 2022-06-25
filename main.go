package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.DELETE("/users", deleteUser)
	router.POST("/users", createUser)
	if err := router.Run("localhost:8080"); err != nil {
		fmt.Println("Error:", err)
	}

	// if err:=db.delete("user","Raghav");err!=nil {
	// 	fmt.Printf("Error while deleting",err)
	// }
}
