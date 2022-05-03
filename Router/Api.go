package Router

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Controller"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Db"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Job"
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
		if port == "" {
			port = "1010"
		}
		//Route init //
		Route = httprouter.New()
		// router group
		Route.GET("/health", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			Helper.ResponseMessage(w, http.StatusOK, "Student Attendance Running Successfully")
		})
		Route.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

			writer.Write([]byte("Welcome to Student Attendance System"))

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
		// Student Attendance  route //
		Route.GET("/student/attendance", Middleware.Auth(Controller.GetRunningStatusForStudent))
		Route.POST("/student/attendance/:id", Middleware.Auth(Controller.SubmitAttendance))
		fmt.Println("Server started on port " + port)
		// conJob
		Job.JobSchedule()
		hand := cors.Default().Handler(Route)
		err := http.ListenAndServe(":"+port, hand)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}

	}
}
