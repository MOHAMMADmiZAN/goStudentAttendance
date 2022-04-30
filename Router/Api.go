package Router

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Controller"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Db"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Middleware"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/config"
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
			config.JobSchedule()
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
		// Admin Attendance  route //
		Route.GET("/admin/attendance", Middleware.Auth(Controller.GetAllAttendance))
		Route.POST("/admin/attendance", Middleware.Auth(Controller.Enable))
		Route.PUT("/admin/attendance/:id", Middleware.Auth(Controller.DisableAttendance))

		fmt.Println("Server started on port " + port)
		// conJob
		config.JobSchedule()
		hand := cors.Default().Handler(Route)
		err := http.ListenAndServe(":"+port, hand)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}

	}
}
