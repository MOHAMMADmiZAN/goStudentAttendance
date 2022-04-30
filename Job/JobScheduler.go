package Job

import (
	"fmt"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Helper"
	"github.com/MOHAMMADmiZAN/goStudentAttendance/Service"
	"github.com/go-co-op/gocron"
	"net/http"
	"time"
)

var TaskArray []func()

func Task1() {
	result := Service.FindActiveAttendance()
	if len(result) > 0 {
		Service.DisableAttendanceWhenTimeOut(result)
		return
	}

	return

}

// JobSchedule JobScheduler is a function that will be called for every scheduled job

func JobSchedule() {
	var w http.ResponseWriter

	TaskArray = append(TaskArray, Task1)

	Schedule := gocron.NewScheduler(time.UTC)

	_, err := Schedule.Every(1).Second().Tag("DisableAttendanceMethod").Do(TaskArray[0])
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

	fmt.Println()

}
