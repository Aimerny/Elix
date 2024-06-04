package onge

import (
	"github.com/aimerny/kook-go/app/core/model"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/client"
	"github/aimerny/elix/app/internal/common"
	"github/aimerny/elix/app/internal/dto"
	"gorm.io/gorm"
)

var (
	OngeServiceDS      *gorm.DB
	OngeStatus         bool
	DeveloperToken     string
	maiChartTypeList   = []string{"Basic", "Advance", "Export", "Master", "Re:Master"}
	chuniChartTypeList = []string{"Basic", "Advance", "Export", "Master", "Ultra", "World Endless"}
)

func RejectOngeProcess(event *model.Event) {
	client.QuotedReplyText("onge service status is disable. notice admin please", event)
}

// InitOngeService Init music game about service
func InitOngeService(conf *common.Config) {
	dbConf := conf.OngeDatasource
	if OngeServiceDS == nil || dbConf != nil {
		_OngeServiceDS, err := dbConf.DbConfig.ConnectDB()
		if err != nil {
			OngeStatus = false
			return
		}
		OngeServiceDS = _OngeServiceDS
		OngeStatus = true
	}
	if conf.DivingFishDeveloperToken != "" {
		DeveloperToken = conf.DivingFishDeveloperToken
	}
	log.Infof("init onge service finished")
	if *common.UpgradeOngeDatabase {
		log.Infof("onge database upgrading ...")
		err := OngeServiceDS.AutoMigrate(
			&dto.MaiMusicInfo{},
			&dto.MaiChartInfo{},
			&dto.ChuniMusicInfo{},
			&dto.ChuniChartInfo{},
			&dto.OngeUserInfo{},
		)
		if err != nil {
			log.WithError(err).Panicf("onge database upgrade failed")
		}
		log.Infof("onge database upgrade success")
	}
}

func BindUser(kookUserId, username string, event *model.Event) {
	user, ok := FindUser(kookUserId)
	if ok && user.DivingUsername != "" && user.DivingUsername != username {
		client.QuotedReplyText("*检测到已绑定到账号<"+user.DivingUsername+">,即将覆盖...*", event)
	}
	user.DivingUsername = username
	user.KookId = kookUserId
	OngeServiceDS.Save(user)
	client.QuotedReplyText("*已绑定到账号:"+username+"*", event)
}

func FindUser(kookUserId string) (*dto.OngeUserInfo, bool) {
	user := &dto.OngeUserInfo{}
	OngeServiceDS.Where(&dto.OngeUserInfo{
		KookId: kookUserId,
	}).Find(user)
	if user.ID != 0 {
		return user, true
	}
	return user, false
}
