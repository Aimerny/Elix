package dto

type DivingMaiResp []*DivingMaiMusicInfo

type DivingMaiMusicInfo struct {
	Id           string                `json:"id"`
	Title        string                `json:"title"`
	Type         string                `json:"type"`
	Difficulties []float32             `json:"ds"`
	Levels       []string              `json:"level"`
	Cids         []int                 `json:"cids"`
	Charts       []*DivingChartInfo    `json:"charts"`
	BasicInfo    *DivingMusicBasicInfo `json:"basic_info"`
}

type DivingChuniResp []*DivingChuniMusicInfo

type DivingChuniMusicInfo struct {
	Id           int                   `json:"id"`
	Title        string                `json:"title"`
	Difficulties []float32             `json:"ds"`
	Levels       []string              `json:"level"`
	Cids         []int                 `json:"cids"`
	Charts       []*DivingChartInfo    `json:"charts"`
	BasicInfo    *DivingMusicBasicInfo `json:"basic_info"`
}

type DivingChartInfo struct {
	Notes   []int  `json:"notes"`
	Combo   int    `json:"combo"`
	Charter string `json:"charter"`
}

type DivingMusicBasicInfo struct {
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`
	BPM         int    `json:"bpm"`
	ReleaseDate string `json:"release_date"`
	From        string `json:"from"`
	IsNew       bool   `json:"is_new"`
}
