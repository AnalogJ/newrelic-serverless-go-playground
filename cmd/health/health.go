package main

import (
	"context"
	"fmt"
	"github.com/analogj/newrelic-serverless-go-playground/pkg"
	"github.com/analogj/newrelic-serverless-go-playground/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
	"log"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	// Create a New Relic application. This will look for your license key in an
	// environment variable called NEW_RELIC_LICENSE_KEY. This example turns on
	// Distributed Tracing, but that's not required.
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDebugLogger(os.Stdout),
		nrlambda.ConfigOption(),
		func(config *newrelic.Config) {
			logrus.SetLevel(logrus.DebugLevel)
			config.Logger = nrlogrus.StandardLogger()
		},
	)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	nrlambda.Start(handlers.HandlerMiddleware(
		func(ctx context.Context, request events.APIGatewayV2HTTPRequest, appLogger *logrus.Entry) (interface{}, pkg.ResponseType, error) {
			txn := newrelic.FromContext(ctx)
			txn.AddAttribute("userLevel", "gold")
			txn.Application().RecordCustomEvent("MyEvent", map[string]interface{}{
				"zip": "zap",
			})

			log.Printf("logrus logger hello world from log function")
			fmt.Printf("logrus logger  hello world from fmt function")
			appLogger.Printf("logrus logger hello world from appLogger function")
			//health check is always true at this point
			return true, pkg.ResponseTypeJSON, nil
		},
	), app)
}
