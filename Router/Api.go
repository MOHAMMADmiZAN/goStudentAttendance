package Router

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Controller"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Db"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helpers"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Middleware"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

var Route *httprouter.Router

func Api() {
	Helpers.LoadEnv()
	// portNumber //
	port := os.Getenv("PORT_NUMBER")
	//Route init //
	Route = httprouter.New()
	// user route //
	Route.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Helpers.MyArrayMap(Service.UserRoles, func(v string, i int, arr []string) {
			fmt.Fprintf(w, "%s\n", v)
			fmt.Println(i, v)

		})

	})
	// register route //
	Route.POST("/register", Controller.Register)
	// login route //
	Route.POST("/login", Controller.Login)
	// user route //

	Route.GET("/users", Middleware.Auth(Controller.GetAllUsers))
	Route.POST("/users", Middleware.Auth(Controller.CreateNewUser))
	Route.GET("/users/:id", Middleware.Auth(Controller.GetUser))
	Route.PUT("/users/:id", Middleware.Auth(Controller.UpdateUser))
	Route.DELETE("/users/:id", Middleware.Auth(Controller.DeleteUser))
	// student route //

	// Profile route //

	Db.Init()
	fmt.Println("Server started on port " + port)
	hand := cors.Default().Handler(Route)
	err := http.ListenAndServe(":"+port, hand)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
