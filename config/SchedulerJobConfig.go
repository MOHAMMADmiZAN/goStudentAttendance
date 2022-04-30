package config

import (
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/go-co-op/gocron"
	"net/http"
	"time"
)

func Task1() {
	result := Service.FindActiveAttendance()
	if len(result) > 0 {
		Service.DisableAttendanceWhenTimeOut(result)
		return
	}
	//fmt.Println("No active attendance")
	//return

}

func JobSchedule() {
	var w http.ResponseWriter

	var TaskArray []func()
	TaskArray = append(TaskArray, Task1)

	Schedule := gocron.NewScheduler(time.UTC)

	_, err := Schedule.Every(1).Second().Do(TaskArray[0])
	if err != nil {
		Helper.ResponseMessage(w, http.StatusFailedDependency, err.Error())
	}
	location, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		Helper.ResponseMessage(w, http.StatusFailedDependency, err.Error())
		return
	}
	Schedule.ChangeLocation(location)
	Schedule.StartAsync()

}
