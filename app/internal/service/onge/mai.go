package onge

import (
	"fmt"
	"github.com/aimerny/kook-go/app/core/helper"
	"github.com/aimerny/kook-go/app/core/model"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/dto"
	"gorm.io/gorm"
	"io"
	"strconv"
)

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

func calcMaiChartType(index int) string {
	return maiChartTypeList[index]
}

// FindMaiMusicInfo priority of search: id > music name
func FindMaiMusicInfo(keyword string) *dto.MaiMusicInfo {
	musicInfo := &dto.MaiMusicInfo{}
	result := OngeServiceDS.Where("id = ?", keyword).First(musicInfo)
	if result.RowsAffected == 0 {
		result = OngeServiceDS.Where("title like ?", keyword).First(musicInfo)
	}
	if result.RowsAffected == 0 {
		return nil
	} else {
		return musicInfo
	}
}

func GenMusicCard(music *dto.MaiMusicInfo) string {
	card := model.NewCard(model.ThemeTypeInfo, model.SizeLg)
	modules := make([]model.CardModule, 0)
	//Header
	modules = append(modules, *model.NewTextHeader("你要找的歌是否是:"))
	modules = append(modules, *model.NewImage(fmt.Sprintf("https://www.diving-fish.com/covers/%05d.png", music.ID)))
	modules = append(modules, *model.NewKMarkdown(
		fmt.Sprintf("**%d.(font)%s(font)[danger]**\n**曲师** : %s\n**BPM** : %d", music.ID, music.Title, music.Artist, music.Bpm)),
	)
	card.Modules = modules
	req := []*model.CardModule{card}
	data, err := jsoniter.Marshal(req)
	if err != nil {
		log.WithError(err).Error("gen music info failed")
		return ""
	}
	return string(data)
}
