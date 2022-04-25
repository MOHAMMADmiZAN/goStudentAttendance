package Db

import (
	"fmt"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Init() {
	err := mgm.SetDefaultConfig(nil, "go-student-attendance", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Database connection error: ", err)
	} else {
		fmt.Println("Database connection successful")
	}

}
