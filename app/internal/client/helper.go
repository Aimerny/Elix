package client

//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/tencent-connect/botgo/log"
//	"io"
//	"net/http"
//	"time"
//)
//
////const QQApiUrl = "https://api.sgroup.qq.com"
//
//type QQClient struct {
//	http.Client
//	meta        *BotMeta
//	accessToken string
//}
//
//type BotMeta struct {
//	AppId     string `json:"appId"`
//	AppSecret string `json:"clientSecret"`
//}
//
//type KookClient struct {
//	http.Client
//	AccessToken string
//}
//
//func NewQQClient(appId string, accessToken string) *QQClient {
//	client := &QQClient{
//		http.Client{},
//		&BotMeta{
//			appId,
//			accessToken,
//		},
//		"",
//	}
//	go client.flushToken()
//
//	return client
//}
//
//func (c *QQClient) Get(url string) ([]byte, error) {
//
//	request, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		log.Errorf("Http request construct failed! err: %e", err)
//		return nil, err
//	}
//	request.Header.Set("Authorization", fmt.Sprintf("QQBot %s", c.accessToken))
//	request.Header.Set("X-Union-Appid", c.meta.AppId)
//	return c.doRequestAndReadBody(request)
//}
//
//func (c *QQClient) Post(url string, body any) ([]byte, error) {
//	bodyBytes, _ := json.Marshal(body)
//	buffer := bytes.NewBuffer(bodyBytes)
//	request, err := http.NewRequest("POST", url, buffer)
//	if err != nil {
//		log.Errorf("Http request construct failed! err: %e", err)
//		return nil, err
//	}
//	request.Header.Set("Authorization", fmt.Sprintf("QQBot %s", c.accessToken))
//	request.Header.Set("X-Union-Appid", c.meta.AppId)
//	return c.doRequestAndReadBody(request)
//}
//
//func (c *QQClient) doRequestAndReadBody(r *http.Request) ([]byte, error) {
//	resp, err := c.Do(r)
//	if err != nil {
//		log.Errorf("http request failed! err:%e", err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//	data, err := io.ReadAll(resp.Body)
//	if err != nil {
//		log.Errorf("read response body failed! err:%e", err)
//		return nil, err
//	}
//	return data, nil
//}
//
//func (c *QQClient) flushToken() {
//	// https://bots.qq.com/app/getAppAccessToken
//	log.Info("start flush token")
//	body, _ := json.Marshal(c.meta)
//	buffer := bytes.NewBuffer(body)
//	request, _ := http.NewRequest("POST", "https://bots.qq.com/app/getAppAccessToken", buffer)
//	request.Header.Add("Content-Type", "application/json")
//	for {
//		// 1min 尝试刷新一次
//		resp, err := c.Do(request)
//		if err != nil {
//			log.Errorf("[QQClient] get access token failed")
//			continue
//		}
//		defer resp.Body.Close()
//		authInfoBytes, _ := io.ReadAll(resp.Body)
//		authInfo := make(map[string]string)
//		err = json.Unmarshal(authInfoBytes, &authInfo)
//		if err != nil {
//			log.Errorf("[QQClient] unmarshal resp body failed. resp: [%s]", string(authInfoBytes))
//			continue
//		}
//		if value := authInfo["access_token"]; value != "" {
//			if c.accessToken != value {
//				log.Info("[QQClient] accessToken has flushed, new token is:", value)
//				c.accessToken = authInfo["access_token"]
//			}
//		}
//		<-time.After(1 * time.Minute)
//	}
//}
