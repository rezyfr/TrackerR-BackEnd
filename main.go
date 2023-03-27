package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

var users = []User{
	{Email: "frizkillah98@gmail.com", Name: "Fidriyanto Rizkillah"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByEmail(c *gin.Context) {
	email := c.Param("email")
	for _, a := range users {
		if a.Email == email {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func postUsers(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	host, errEnv := os.LookupEnv("dbhost")
	log.Println(host, errEnv)

	dbport, errEnv := os.LookupEnv("dbport")
	log.Println(host, errEnv)

	const (
		user     = "postgres"
		password = "MXUb%$JV}AL-,2[j"
		dbname   = "postgres"
		sslmode  = "disable"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, dbport, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:email", getUserByEmail)
	router.POST("/users", postUsers)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("0.0.0.0:" + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
