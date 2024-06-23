package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/josgarc0/gambit/awsgo"
	"github.com/josgarc0/gambit/bd"
	"github.com/josgarc0/gambit/handlers"
	/*
		"github.com/aws/aws-lambda-go/events"
		lambda "github.com/aws/aws-lambda-go/lambda"
		"github.com/josgarc0/gambit/awsgo"
		"github.com/josgarc0/gambit/bd"
	*/)

func main() {
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		//fmt.Println("Error en los parametros. Debe enviar 'Secret Manager','UrlPrefix'")
		panic("Error en los parametros. Debe enviar 'Secret Manager','UrlPrefix'")
		//err := errors.New("error en los parametros, debe enviar SecretName")
		//return event, err
	}

	var res *events.APIGatewayProxyResponse
	//prefix := os.Getenv("UrlPrefix")
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)

	method := request.RequestContext.HTTP.Method

	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, header, request)
	//
	headersResp := map[string]string{
		"content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil
}

func ValidoParametros() bool {

	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if traeParametro {
		return traeParametro
	}
	return true
}
