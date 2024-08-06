package onge

import (
	"github.com/aimerny/kook-go/app/core/helper"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/dto"
	"io"
)

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
				Model:      dto.Model{ID: uint(divingMusicInfo.Cids[5])},
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
					Model:      dto.Model{ID: uint(divingMusicInfo.Cids[index])},
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

func calcChuniChartType(index int) string {
	return chuniChartTypeList[index]
}
