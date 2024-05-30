package server

import (
	"github.com/aimerny/kook-go/app/core/action"
	"github.com/aimerny/kook-go/app/core/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/service"
	"io"
	"net/http"
	"strconv"
)

func StartApiServer(port int) {
	http.HandleFunc("/message/send", messageSend)
	http.HandleFunc("/channel/bot-channels", getAllBotChannelsMeta)
	logrus.Info("start kook api server")
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		logrus.WithError(err).Error("api server listen stop")
		return
	}
}

func messageSend(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		logrus.WithError(err).Error("server read req failed")
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	msgReq := &model.MessageCreateReq{}
	err = jsoniter.Unmarshal(bodyBytes, msgReq)
	if err != nil {
		logrus.WithError(err).Error("unmarshal req body failed")
		resp.WriteHeader(http.StatusBadRequest)
	}
	action.MessageSend(msgReq)
	logrus.WithField("msg", msgReq).Info("server send req")
}

func getAllBotChannelsMeta(resp http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	searchKey := query.Get("searchKey")
	channels := service.FindChannels(searchKey)
	bytes, _ := jsoniter.Marshal(channels)
	resp.WriteHeader(http.StatusOK)
	resp.Write(bytes)
}
