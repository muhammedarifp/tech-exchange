package cronejobs

import (
	"github.com/robfig/cron"
)

func InitCronJobs() {
	c := cron.New()
	exp := "*/10 * * * * *"
	c.AddFunc(exp, awardOneYearBadge)
	c.Start()
}
