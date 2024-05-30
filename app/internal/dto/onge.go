package dto

import "gorm.io/gorm"

type MaiMusicInfo struct {
	gorm.Model
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
	gorm.Model
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
	gorm.Model
	Title       string
	Bpm         int
	Artist      string
	Genre       string
	From        string
	ReleaseDate string
}

type ChuniChartInfo struct {
	gorm.Model
	MusicId    int
	Type       string
	Difficulty float32
	Level      string
	Charter    string
	Combo      int
}
