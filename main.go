package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vimal/mongo-golang/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()

	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	log.Fatal(http.ListenAndServe("localhost:9090", r))
	fmt.Println("Connected localhost:9090")
}

func getSession() *mgo.Session {

	// connect with mongo db
	s, err := mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	return s
}
