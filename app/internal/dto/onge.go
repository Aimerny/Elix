package dto

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeSave 在保存前的回调
func (model *Model) BeforeSave(tx *gorm.DB) (err error) {
	if model.ID != 0 {
		var existingModel Model
		if err := tx.Unscoped().Table(tx.Statement.Table).First(&existingModel, model.ID).Error; err == nil {
			model.CreatedAt = existingModel.CreatedAt
		}
	} else {
		model.CreatedAt = time.Now()
	}
	model.UpdatedAt = time.Now()
	return
}

type MaiMusicInfo struct {
	Model
	Title       string
	Type        string
	Bpm         int
	Artist      string
	Genre       string
	From        string
	ReleaseDate string
	IsNew       bool
}

type MaiChartInfo struct {
	Model
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	MusicId    int
	Type       string
	Difficulty float32
	Level      string
	Charter    string

	TapNote   int
	HoldNote  int
	SlideNote int
	TouchNote int
	BreakNote int
	Combo     int
}

type ChuniMusicInfo struct {
	Model
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Bpm         int
	Artist      string
	Genre       string
	From        string
	ReleaseDate string
}

type ChuniChartInfo struct {
	Model
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	MusicId    int
	Type       string
	Difficulty float32
	Level      string
	Charter    string
	Combo      int
}

type OngeUserInfo struct {
	Model
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	KookId          string
	DivingFishToken string
	DivingUsername  string
}
