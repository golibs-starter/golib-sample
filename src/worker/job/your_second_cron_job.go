package job

import (
	"context"
	golibcron "github.com/golibs-starter/golib-cron"
	"github.com/golibs-starter/golib/log"
)

type YourSecondCronJob struct {
}

func (y YourSecondCronJob) Name() string {
	return "YourCustomizedSecondCronJobName"
}

func NewYourSecondCronJob() golibcron.Job {
	return &YourSecondCronJob{}
}

func (y YourSecondCronJob) Run(ctx context.Context) {
	log.Infoc(ctx, "YourSecondCronJob started")
}
