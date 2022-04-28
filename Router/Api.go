package Router

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Controller"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Db"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

var Route *httprouter.Router

func Api() {
	if Db.Init() {
		// portNumber //
		port := os.Getenv("PORT_NUMBER")
		//Route init //
		Route = httprouter.New()
		// user route //
		Route.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			fmt.Println(Helper.RandomString(32))
		})
		// Auth route //
		Route.POST("/register", Controller.Register)

		Route.POST("/login", Controller.Login)

		// user route //
		Route.GET("/users", Middleware.Auth(Controller.GetAllUsers))
		Route.POST("/users", Middleware.Auth(Controller.CreateNewUser))
		Route.GET("/users/:id", Middleware.Auth(Controller.GetUser))
		Route.PUT("/users/:id", Middleware.Auth(Controller.UpdateUser))
		Route.DELETE("/users/:id", Middleware.Auth(Controller.DeleteUser))
		// student route //

		fmt.Println("Server started on port " + port)
		hand := cors.Default().Handler(Route)
		err := http.ListenAndServe(":"+port, hand)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}

	}
}
