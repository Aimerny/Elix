package client

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/dto"
	"io"
	"net/http"
	"net/url"
)

const (
	DivingFishPlayerScoreQuery = "https://www.diving-fish.com/api/maimaidxprober/dev/player/records"
)

func Get(req *http.Request) ([]byte, error) {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithError(err).Error("do get failed! err")
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("action read body failed!")
		return nil, err
	}
	return data, nil
}

func QueryRecord(username string, devToken string) (*dto.DivingPlayerRecordsResp, error) {

	reqUrl, _ := url.Parse(DivingFishPlayerScoreQuery)
	query := reqUrl.Query()
	query.Set("username", username)
	reqUrl.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		log.Errorf("Http request construct failed! err: %e", err)
		return nil, err
	}
	req.Header.Add("Developer-Token", devToken)
	body, err := Get(req)
	resp := &dto.DivingPlayerRecordsResp{}
	err = jsoniter.Unmarshal(body, resp)
	if err != nil {
		log.WithError(err).WithField("resp", string(body)).Error("unmarshal record failed")
		return nil, err
	}
	if resp.Status == "error" {
		err := errors.New(resp.Message)
		log.WithError(err).WithField("resp", string(body)).Error("query record failed")
		return nil, err
	}
	return resp, nil
}
