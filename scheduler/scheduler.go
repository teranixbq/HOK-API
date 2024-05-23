package scheduler

import (
	"fmt"
	"hokapi/app/middleware"
	"time"
	"github.com/go-co-op/gocron/v2"
)

func ResetRequestCountAtMidnight() {
	location, _ := time.LoadLocation("Asia/Jakarta")
	next := location
	timeGMT7 := time.Now().In(next)

	if timeGMT7.Hour() == 0 {
		if middleware.RequestCount > 1000 {
			middleware.RequestCount = 0
			fmt.Println("Request count has been reset")
		}
	}
}

func Scheduler() error {

	s, err := gocron.NewScheduler()
	if err != nil {
		return err
	}

	_, err = s.NewJob(
		gocron.DurationJob(
			20*time.Second,
		),
		gocron.NewTask(
			func() {
				ResetRequestCountAtMidnight()
			},
		),
	)

	if err != nil {
		return err
	}

	s.Start()
	
	return nil
}