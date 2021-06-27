package util

import (
	"log"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func init() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              AppConfig.SentryDSN,
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalln("Could not initialize Sentry: ", err)
	}
	if AppConfig.SentryDSN == "" {
		log.Println("Sentry is disabled")
	} else {
		log.Println("Initialized Sentry")
	}
}

func NotifyError(ctx *gin.Context, err error) {
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}
}
