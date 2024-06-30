package pns

import (
	"errors"
	"log"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dypnsapi "github.com/alibabacloud-go/dypnsapi-20170525/client"
)

type PNS struct {
	client *dypnsapi.Client
}

func New(accessKeyID string, accessKeySecret string, endpoint string) *PNS {
	cfg := &openapi.Config{
		AccessKeyId:     &accessKeyID,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
	}

	cli, err := dypnsapi.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &PNS{
		client: cli,
	}
}

func (u *PNS) GetMobile(token string) (*dypnsapi.GetMobileResponse, error) {
	req := &dypnsapi.GetMobileRequest{
		AccessToken: &token,
	}
	resp, err := u.client.GetMobile(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if *resp.Body.Code != "OK" {
		return nil, errors.New("code error")
	}
	return resp, nil
}
