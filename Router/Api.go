package Router

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Db"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func Api() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// portNumber //
	port := os.Getenv("PORT_NUMBER")
	//Route init //
	Route := httprouter.New()

	Db.Init()
	err = http.ListenAndServe(port, Route)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
