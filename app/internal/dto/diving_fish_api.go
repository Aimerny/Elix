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

type DivingPlayerRecordsResp struct {
	Username         string                   `json:"username,omitempty"`
	AdditionalRating int                      `json:"additional_rating,omitempty"`
	Nickname         string                   `json:"nickname,omitempty"`
	Plate            string                   `json:"plate,omitempty"`
	Rating           int                      `json:"rating"`
	Records          []DivingPlayerRecordInfo `json:"records"`
	Status           string                   `json:"status"`
	Message          string                   `json:"msg"`
}

type DivingPlayerB50Info struct {
	DivingPlayerRecordsResp
	B35 *[]*DivingPlayerRecordInfo `json:"b35"`
	B15 *[]*DivingPlayerRecordInfo `json:"b15"`
}

type DivingPlayerRecordInfo struct {
	Achievements float64 `json:"achievements,omitempty"`
	Ds           float64 `json:"ds,omitempty"`
	DxScore      int     `json:"dx_score,omitempty"`
	Fc           string  `json:"fc"`
	Fs           string  `json:"fs"`
	Level        string  `json:"level"`
	LevelIndex   int     `json:"level_index"`
	LevelLabel   string  `json:"level_label"`
	Ra           int     `json:"ra"`
	Rate         string  `json:"rate"`
	SongID       uint    `json:"song_id"`
	Title        string  `json:"title"`
	Type         string  `json:"type"`
	Cover        string  `json:"cover,omitempty"`
}
