package Router

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Controller"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Db"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

var Route *httprouter.Router

func Api() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// portNumber //
	port := os.Getenv("PORT_NUMBER")
	//Route init //
	Route = httprouter.New()
	// user route //
	Route.POST("/user", Controller.CreateUser)
	Db.Init()
	fmt.Println("Server started on port " + port)
	hand := cors.Default().Handler(Route)
	err = http.ListenAndServe(":"+port, hand)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
