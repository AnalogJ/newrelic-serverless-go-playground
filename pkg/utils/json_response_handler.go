package utils

import (
	"bytes"
	"encoding/json"
	"github.com/analogj/newrelic-serverless-go-playground/pkg/models"
	"github.com/aws/aws-lambda-go/events"
)

func JsonResponseHandler(bodyWrapper models.ResponseWrapper, unwrapped bool) (events.APIGatewayV2HTTPResponse, error) {
	var buf bytes.Buffer
	response := events.APIGatewayV2HTTPResponse{
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
			//"Access-Control-Allow-Origin":      "*",
			//"Access-Control-Allow-Credentials": "true",
		},
	}
	if bodyWrapper.Success {
		response.StatusCode = 200
	} else {
		response.StatusCode = 500
	}
	var err error
	var body []byte
	if unwrapped {
		body, err = json.Marshal(bodyWrapper.Data)
		json.HTMLEscape(&buf, body)
	} else {
		body, err = json.Marshal(bodyWrapper)
		json.HTMLEscape(&buf, body)
	}

	response.Body = buf.String()
	return response, err
}
