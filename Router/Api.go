package Router

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Controller"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Db"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Model"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/kamva/mgm/v3"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
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
	Route.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		email := "mizan@gmail.com"
		user := &Model.User{}
		coll := mgm.Coll(user)
		err = coll.First(bson.M{"email": email}, user)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)

	})
	Route.POST("/user", Controller.CreateUser)

	Db.Init()
	fmt.Println("Server started on port " + port)
	hand := cors.Default().Handler(Route)
	err = http.ListenAndServe(":"+port, hand)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
