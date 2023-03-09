package main

import (
	"github.com/andrsj/feedback-service/internal/app"
	log "github.com/andrsj/feedback-service/pkg/logger"
	zap "github.com/andrsj/feedback-service/pkg/logger/zap"
)

func main() {
	zap := zap.New()

	app, err := app.New(zap)
	if err != nil {
		zap.Fatal("Can't configure the app", log.M{"err": err})
	}

	err = app.Start()
	if err != nil {
		zap.Fatal("Can't start the application", log.M{"err": err})
	}
}
