package service

import (
	"github.com/aimerny/kook-go/app/core/helper"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/common"
	"github/aimerny/elix/app/internal/dto"
	"gorm.io/gorm"
	"io"
	"strconv"
)

var OngeServiceDS *gorm.DB
var maiChartTypeList = []string{"Basic", "Advance", "Export", "Master", "Re:Master"}
var chuniChartTypeList = []string{"Basic", "Advance", "Export", "Master", "Ultra", "World Endless"}

// InitOngeService Init music game about service
func InitOngeService(dbConf *common.DatasourceConf) {
	if OngeServiceDS != nil || dbConf != nil {
		OngeServiceDS = dbConf.DbConfig.ConnectDB()
	}
	log.Infof("init onge service finished")
}

func FlushMaimaiDB() {
	maimaiResp, err := helper.Get("https://www.diving-fish.com/api/maimaidxprober/music_data")
	if err != nil {
		log.WithError(err).Error("get maimai meta failed")
		return
	}
	defer maimaiResp.Body.Close()
	data, err := io.ReadAll(maimaiResp.Body)
	if err != nil {
		log.Errorf("action read body failed! err:%e", err)
		return
	}
	var maiRespBody dto.DivingMaiResp
	err = jsoniter.Unmarshal(data, &maiRespBody)
	if err != nil {
		log.WithError(err).Error("unmarshal json failed")
	}

	OngeServiceDS.AutoMigrate(dto.MaiMusicInfo{}, dto.MaiChartInfo{})
	for _, divingMusicInfo := range maiRespBody {
		// music
		newMusic := &dto.MaiMusicInfo{}
		idInt, _ := strconv.Atoi(divingMusicInfo.Id)
		newMusic.ID = uint(idInt)
		newMusic.Type = divingMusicInfo.Type
		newMusic.Title = divingMusicInfo.Title
		newMusic.Genre = divingMusicInfo.BasicInfo.Genre
		newMusic.IsNew = divingMusicInfo.BasicInfo.IsNew
		newMusic.Artist = divingMusicInfo.BasicInfo.Artist
		newMusic.Bpm = divingMusicInfo.BasicInfo.BPM
		newMusic.ReleaseDate = divingMusicInfo.BasicInfo.ReleaseDate
		newMusic.From = divingMusicInfo.BasicInfo.From

		OngeServiceDS.Save(newMusic)

		// charts
		for index, chart := range divingMusicInfo.Charts {

			newChart := &dto.MaiChartInfo{
				Model:      gorm.Model{ID: uint(divingMusicInfo.Cids[index])},
				MusicId:    idInt,
				Type:       calcMaiChartType(index),
				Difficulty: divingMusicInfo.Difficulties[index],
				Level:      divingMusicInfo.Levels[index],
				Charter:    chart.Charter,
				TapNote:    chart.Notes[0],
				HoldNote:   chart.Notes[1],
				SlideNote:  chart.Notes[2],
			}
			if len(chart.Notes) > 4 {
				newChart.TouchNote = chart.Notes[3]
				newChart.BreakNote = chart.Notes[4]
				newChart.Combo = chart.Notes[0] + chart.Notes[1] + chart.Notes[2] + chart.Notes[3] + chart.Notes[4]
			} else {
				// old chart has not touch
				newChart.TouchNote = 0
				newChart.BreakNote = chart.Notes[3]
				newChart.Combo = chart.Notes[0] + chart.Notes[1] + chart.Notes[2] + chart.Notes[3]
			}
			OngeServiceDS.Save(newChart)
		}
	}
	return
}

func FlushChuniDB() {
	chuniResp, err := helper.Get("https://www.diving-fish.com/api/chunithmprober/music_data")
	if err != nil {
		log.WithError(err).Error("get chuni meta failed")
		return
	}
	defer chuniResp.Body.Close()
	data, err := io.ReadAll(chuniResp.Body)
	if err != nil {
		log.Errorf("action read body failed! err:%e", err)
		return
	}
	var chuniRespBody dto.DivingChuniResp
	err = jsoniter.Unmarshal(data, &chuniRespBody)
	if err != nil {
		log.WithError(err).Error("unmarshal json failed")
		return
	}

	OngeServiceDS.AutoMigrate(dto.ChuniMusicInfo{}, dto.ChuniChartInfo{})
	for _, divingMusicInfo := range chuniRespBody {
		// music
		newMusic := &dto.ChuniMusicInfo{}
		newMusic.ID = uint(divingMusicInfo.Id)
		newMusic.Title = divingMusicInfo.Title
		newMusic.Artist = divingMusicInfo.BasicInfo.Artist
		newMusic.Bpm = divingMusicInfo.BasicInfo.BPM
		newMusic.ReleaseDate = divingMusicInfo.BasicInfo.ReleaseDate
		newMusic.Genre = divingMusicInfo.BasicInfo.Genre
		newMusic.From = divingMusicInfo.BasicInfo.From

		OngeServiceDS.Save(newMusic)
		// check world end music, only one chart
		if len(divingMusicInfo.Charts) > 5 {
			newChart := &dto.ChuniChartInfo{
				Model:      gorm.Model{ID: uint(divingMusicInfo.Cids[5])},
				MusicId:    divingMusicInfo.Id,
				Type:       calcChuniChartType(5),
				Difficulty: divingMusicInfo.Difficulties[5],
				Level:      divingMusicInfo.Levels[5],
				Charter:    divingMusicInfo.Charts[5].Charter,
				Combo:      divingMusicInfo.Charts[5].Combo,
			}
			OngeServiceDS.Save(newChart)
		} else {
			// normal charts
			for index, chart := range divingMusicInfo.Charts {
				newChart := &dto.ChuniChartInfo{
					Model:      gorm.Model{ID: uint(divingMusicInfo.Cids[index])},
					MusicId:    divingMusicInfo.Id,
					Type:       calcChuniChartType(index),
					Difficulty: divingMusicInfo.Difficulties[index],
					Level:      divingMusicInfo.Levels[index],
					Charter:    chart.Charter,
					Combo:      chart.Combo,
				}
				OngeServiceDS.Save(newChart)
			}
		}
	}
}

func calcMaiChartType(index int) string {
	return maiChartTypeList[index]
}

func calcChuniChartType(index int) string {
	return chuniChartTypeList[index]
}
