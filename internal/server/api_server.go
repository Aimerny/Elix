package server

import (
	"github.com/aimerny/kook-go/core/action"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

func StartApiServer(port int) {
	http.HandleFunc("/message/send", messageSend)
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
	msgReq := &action.MessageCreateReq{}
	err = jsoniter.Unmarshal(bodyBytes, msgReq)
	if err != nil {
		logrus.WithError(err).Error("unmarshal req body failed")
		resp.WriteHeader(http.StatusBadRequest)
	}
	action.MessageSend(msgReq)
	logrus.WithField("msg", msgReq).Info("server send req")
}
