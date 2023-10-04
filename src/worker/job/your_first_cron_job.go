package job

import (
	"context"
	golibcron "github.com/golibs-starter/golib-cron"
	"github.com/golibs-starter/golib/log"
)

type YourFirstCronJob struct {
}

func NewYourFirstCronJob() golibcron.Job {
	return &YourFirstCronJob{}
}

func (y YourFirstCronJob) Run(ctx context.Context) {
	log.Infoc(ctx, "YourFirstCronJob started")
}
