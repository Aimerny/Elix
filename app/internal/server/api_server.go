package server

import (
	"github.com/aimerny/kook-go/app/core/action"
	"github.com/aimerny/kook-go/app/core/model"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/server/middleware"
	"github/aimerny/elix/app/internal/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func StartApiServer(port int) {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	router := gin.Default()
	router.Use(middleware.LogMiddleware())
	router.POST("/message/send", messageSend)
	router.GET("/channel/bot-channels", getAllBotChannelsMeta)

	templateFiles, err := getHTMLFiles("templates")
	if err != nil {
		logrus.WithError(err).Panic("start api server failed!")
	}
	router.LoadHTMLFiles(templateFiles...)
	router.Static("/assets", "./templates/assets")
	router.GET("/onge/mai/b50", renderMaiB50)
	logrus.Info("start api server")
	err = router.Run(":" + strconv.Itoa(port))
	if err != nil {
		logrus.WithError(err).Error("api server listen stop")
		return
	}
}

func messageSend(ctx *gin.Context) {
	req := ctx.Request
	defer req.Body.Close()
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		logrus.WithError(err).Error("server read req failed")
		ctx.Status(http.StatusBadRequest)
		return
	}
	msgReq := &model.MessageCreateReq{}
	err = jsoniter.Unmarshal(bodyBytes, msgReq)
	if err != nil {
		logrus.WithError(err).Error("unmarshal req body failed")
		ctx.Status(http.StatusBadRequest)
	}
	action.MessageSend(msgReq)
	logrus.WithField("msg", msgReq).Info("server send req")
}

func getAllBotChannelsMeta(ctx *gin.Context) {
	searchKey := ctx.Query("searchKey")
	channels := service.FindChannels(searchKey)
	ctx.JSON(http.StatusOK, channels)
}

func getHTMLFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
