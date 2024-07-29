package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/client"
	"github/aimerny/elix/app/internal/service"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func StartWsProxyServer(port int) {

	http.HandleFunc("/ws", EchoMessage)
	logrus.Info("Ready to start ws server")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", port), nil)
	if err != nil {
		logrus.Error("Failed to start a websocket server")
	}
}

func EchoMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Warning("upgrade failed!")
	}
	//add client to cache
	service.Clients = append(service.Clients, &client.WsClient{
		Conn:   conn,
		Status: true,
	})
	defer conn.Close()
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.Warning("ws read message error, remove the conn from clients:", err)

			break
		}

		//todo other op
		logrus.Infof("%s send to: %s\n", conn.RemoteAddr(), string(msg))
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}

}

func checkOrigin(r *http.Request) bool {
	return true
}
