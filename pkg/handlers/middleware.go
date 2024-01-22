package handlers

import (
	"context"
	"github.com/analogj/newrelic-serverless-go-playground/pkg"
	"github.com/analogj/newrelic-serverless-go-playground/pkg/models"
	"github.com/analogj/newrelic-serverless-go-playground/pkg/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

// this function wraps every standard function, initializing common singletons like Config, Logger and Database
// replicates funtionality middleware in Gin
// also handles errors in a consistent way
// https://www.zachjohnsondev.com/posts/lambda-go-middleware/
func HandlerMiddleware(
	//handler wrapped by this middleware function.
	handler func(
		ctx context.Context,
		request events.APIGatewayV2HTTPRequest,
		appLogger *logrus.Entry,
	) (interface{}, pkg.ResponseType, error)) func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	return func(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		responseData := models.ResponseWrapper{
			Success: true,
		}

		//create logger
		appLogger := logrus.New().WithContext(ctx)

		// call the handler
		data, responseType, err := handler(ctx, request, appLogger)
		if err != nil {
			responseData.Success = false
			responseData.Error = err.Error()
			return utils.JsonResponseHandler(responseData, false)
		} else if responseType == pkg.ResponseTypeJSON {
			responseData.Success = true
			responseData.Data = data
			return utils.JsonResponseHandler(responseData, false)
		} else if responseType == pkg.ResponseTypeRedirect {
			responseData.Success = true
			responseData.Data = data
			return utils.RedirectResponseHandler(responseData)
		} else {
			responseData.Success = false
			responseData.Error = "unknown response type"
			return utils.JsonResponseHandler(responseData, false)
		}

	}
}
