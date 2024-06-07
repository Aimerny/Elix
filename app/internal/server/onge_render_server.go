package server

import (
	"github.com/gin-gonic/gin"
	"github/aimerny/elix/app/internal/dto"
	"github/aimerny/elix/app/internal/service/onge"
	"net/http"
)

func renderMaiB50(ctx *gin.Context) {
	username := "seeyou"
	b50, err := onge.QueryMaiB50(username)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}
	ctx.HTML(http.StatusOK, "b50.html", gin.H{
		"b35":              *splitArrayIntoChunks(b50.B35, 5),
		"b15":              *splitArrayIntoChunks(b50.B15, 5),
		"Username":         b50.Username,
		"AdditionalRating": b50.AdditionalRating,
		"Nickname":         b50.Nickname,
		"Plate":            b50.Plate,
		"Rating":           b50.Rating,
		"Status":           b50.Status,
		"Message":          b50.Message,
	})
}

func splitArrayIntoChunks(arr *[]*dto.DivingPlayerRecordInfo, chunkSize int) *[][]*dto.DivingPlayerRecordInfo {
	var result [][]*dto.DivingPlayerRecordInfo

	for i := 0; i < len(*arr); i += chunkSize {
		end := i + chunkSize

		if end > len(*arr) {
			end = len(*arr)
		}

		result = append(result, (*arr)[i:end])
	}

	return &result
}
