package Db

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helpers"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func Init() {
	Helpers.LoadEnv()
	DBNAME := os.Getenv("DATA_BASE_NAME")
	DBURI := os.Getenv("DATA_BASE_URI")
	err := mgm.SetDefaultConfig(nil, DBNAME, options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal("Database connection error: ", err)
	} else {
		fmt.Println("Database connection successful")
	}

}
