package main

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

var clientGlobal *sdk.Client

func init() {
	aliyunRegionID, ok := os.LookupEnv("aliyunRegionID")
	if !ok {
		panic("aliyunRegionID not set")
	}

	aliyunAccessKey, ok := os.LookupEnv("aliyunAccessKey")
	if !ok {
		panic("aliyunAccessKey not set")
	}

	aliyunSecretKey, ok := os.LookupEnv("aliyunSecretKey")
	if !ok {
		panic("aliyunSecretKey not set")
	}

	clientLocal, err := sdk.NewClientWithAccessKey(aliyunRegionID, aliyunAccessKey, aliyunSecretKey)
	if err != nil {
		panic(err)
	}
	clientGlobal = clientLocal
}

type Event struct {
	ToAddress string `json:"ToAddress"`
	Subject string `json:"Subject"`
	Body string `json:"Body"`
}

func handleRequest(ctx context.Context, event Event) (err error) {
	request := requests.NewCommonRequest()
	request.Domain = "dm.aliyuncs.com"
	request.Version = "2015-11-23"
	request.ApiName = "SingleSendMail"

	request.QueryParams["AccountName"] = "no-reply@bytecare.xyz"
	request.QueryParams["AddressType"] = "1"
	request.QueryParams["ReplyToAddress"] = "false"
	request.QueryParams["ToAddress"] = event.ToAddress
	request.QueryParams["Subject"] = event.Subject

	if IsAllWhiteChar(event.Body) {
		request.QueryParams["HtmlBody"] = "<html></html>"
	} else {
		request.QueryParams["TextBody"] = event.Body
	}

	_, err = clientGlobal.ProcessCommonRequest(request)
	return
}

func main() {
	runtime.Start(handleRequest)
}

func IsAllWhiteChar(s string) bool {
	r := strings.TrimSpace(s) == ""
	return r
}