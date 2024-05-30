package command

import (
	"github.com/aimerny/kook-go/app/common"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestParse(t *testing.T) {
	common.InitLogger()
	c, err := Parse("/maimai test 1 1")
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("%v", c)

}
