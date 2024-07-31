package onge

import (
	"fmt"
	"github.com/aimerny/kook-go/app/core/helper"
	"github.com/aimerny/kook-go/app/core/model"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github/aimerny/elix/app/internal/client"
	"github/aimerny/elix/app/internal/common"
	"github/aimerny/elix/app/internal/dto"
	"github/aimerny/elix/app/util"
	"io"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"sync"
)

var NewMaiMusicIds []uint
var conf *common.Config

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
				Model:      dto.Model{ID: uint(divingMusicInfo.Cids[index])},
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

func QueryMaiB50(divingUsername string) (*dto.DivingPlayerB50Info, error) {
	record, err := client.QueryRecord(divingUsername, DeveloperToken)
	if err != nil {
		return nil, err
	}
	//card := model.NewCard(model.ThemeTypeInfo, model.SizeLg)
	//modules := make([]model.CardModule, 0)
	//modules = append(modules, *model.NewKMarkdown(fmt.Sprintf("**Rating:%d**", record.Rating)))
	// best 35
	best35 := make([]*dto.DivingPlayerRecordInfo, 35)
	best15 := make([]*dto.DivingPlayerRecordInfo, 15)

	// 根据年龄排序
	sort.Slice(record.Records, func(i, j int) bool {
		return record.Records[i].Ra > record.Records[j].Ra
	})

	b35Index, b15Index := 0, 0

	for _, recordInfo := range record.Records {
		b15done := b15Index >= len(best15)
		b35done := b35Index >= len(best35)
		if slices.Contains(NewMaiMusicIds, recordInfo.SongID) {
			// calculate in b15
			if !b15done {
				best15[b15Index] = &recordInfo
				recordInfo.Cover = fmt.Sprintf("https://www.diving-fish.com/covers/%05d.png", recordInfo.SongID)
				b15Index++
			}
		} else {
			// calculate in b35
			if !b35done {
				best35[b35Index] = &recordInfo
				recordInfo.Cover = fmt.Sprintf("https://www.diving-fish.com/covers/%05d.png", recordInfo.SongID)
				b35Index++
			}
		}
		if b15done && b35done {
			break
		}
	}
	return &dto.DivingPlayerB50Info{
		DivingPlayerRecordsResp: *record,
		B35:                     &best35,
		B15:                     &best15,
	}, nil
}

// FetchMaiResources get maimai resources
func FetchMaiResources() error {
	conf = common.GlobalConf()
	maiPath := conf.DataDirPath + "/mai"

	// covers
	coverPath := maiPath + "/cover"
	coverDir, err := util.ReadDirForce(coverPath)
	if err != nil {
		return err
	}
	defer coverDir.Close()
	getMusicCovers(coverDir)

	return nil
}

func getMusicCovers(coverDir *os.File) error {
	log.Info("start download maimai music covers")
	defer log.Info("download maimai music covers done!")
	var wg sync.WaitGroup

	info := &dto.MaiMusicInfo{}
	var allMusics []*dto.MaiMusicInfo
	OngeServiceDS.Where(info).Find(&allMusics)
	names, err := coverDir.Readdirnames(-1)
	if err != nil {
		return err
	}
	// concurrent download covers
	semaphore := make(chan struct{}, 5)
	for _, musicInfo := range allMusics {
		wg.Add(1)
		go func(musicId uint, musicName string) {
			semaphore <- struct{}{}
			filename := fmt.Sprintf("%05d.png", musicId)
			if slices.Contains(names, filename) {
				log.WithField("title", musicName).WithField("cover", filename).Debug("found music cover, skip!")
				wg.Done()
			} else {
				downloadMaiCover(musicId, musicName, &wg)
			}
			<-semaphore
		}(musicInfo.Model.ID, musicInfo.Title)
	}
	wg.Wait()
	return nil
}

func downloadMaiCover(musicId uint, musicName string, wg *sync.WaitGroup) {
	defer wg.Done()
	filename := fmt.Sprintf(conf.DataDirPath+"/mai/cover/%05d.png", musicId)
	// 创建文件
	out, err := os.Create(filename)
	if err != nil {
		log.WithError(err).Errorf("Failed to create file: %s", filename)
		return
	}
	defer out.Close()

	// 下载文件
	resp, err := http.Get(fmt.Sprintf("https://www.diving-fish.com/covers/%05d.png", musicId))
	if err != nil {
		log.WithError(err).Errorf("Failed to download cover: %s", filename)
		return
	}
	defer resp.Body.Close()

	// 将内容写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.WithError(err).Errorf("Failed to write cover: %s", filename)
		return
	}
	log.WithField("title", musicName).WithField("cover", filename).Info("Download cover success")
}
