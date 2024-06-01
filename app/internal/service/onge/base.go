package onge

import (
	"github.com/aimerny/kook-go/app/core/model"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/client"
	"github/aimerny/elix/app/internal/common"
	"gorm.io/gorm"
)

var (
	OngeServiceDS      *gorm.DB
	OngeStatus         bool
	maiChartTypeList   = []string{"Basic", "Advance", "Export", "Master", "Re:Master"}
	chuniChartTypeList = []string{"Basic", "Advance", "Export", "Master", "Ultra", "World Endless"}
)

func RejectOngeProcess(event *model.Event) {
	client.QuotedReplyText("onge service status is disable. notice admin please", event)
}

// InitOngeService Init music game about service
func InitOngeService(dbConf *common.DatasourceConf) {
	if OngeServiceDS == nil || dbConf != nil {
		_OngeServiceDS, err := dbConf.DbConfig.ConnectDB()
		if err != nil {
			OngeStatus = false
			return
		}
		OngeServiceDS = _OngeServiceDS
		OngeStatus = true
	}
	log.Infof("init onge service finished")
}
