package utils

import (
	"github.com/analogj/newrelic-serverless-go-playground/pkg/models"
	"github.com/aws/aws-lambda-go/events"
)

func RedirectResponseHandler(bodyWrapper models.ResponseWrapper) (events.APIGatewayV2HTTPResponse, error) {
	response := events.APIGatewayV2HTTPResponse{
		IsBase64Encoded: false,
		Headers: map[string]string{
			//"Content-Type": "application/json",
			//"Access-Control-Allow-Origin":      "*",
			//"Access-Control-Allow-Credentials": "true",
			"Location": bodyWrapper.Data.(string),
		},
	}
	response.StatusCode = 302
	return response, nil
}
