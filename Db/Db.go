package Db

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func Init() bool {
	Helper.LoadEnv()
	DBNAME := os.Getenv("DATA_BASE_NAME")
	DBURI := os.Getenv("DATA_BASE_URI")
	err := mgm.SetDefaultConfig(nil, DBNAME, options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal("Database connection error: ", err)
		return false

	}
	fmt.Println("Database connection successful")
	return true

}
