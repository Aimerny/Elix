package common

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TypeMysql DatasourceType = "mysql"
	// TypeSqlite DatasourceType = "sqlite"
)

type DatasourceType string

type DatasourceConf struct {
	Type      DatasourceType `json:"type"`
	SourceKey string         `json:"source_key"`
	DbConfig  DbConfig       `json:"config"`
}

type _DatasourceConfig DatasourceConf

func (c *DatasourceConf) UnmarshalJSON(b []byte) error {
	d := _DatasourceConfig{}
	switch DatasourceType(jsoniter.Get(b, "type").ToString()) {
	case TypeMysql:
		d.DbConfig = &MysqlDBConfig{}
	default:
		return errors.New("unsupported ds type")
	}
	if err := jsoniter.Unmarshal(b, &d); err != nil {
		return err
	}
	*c = (DatasourceConf)(d)
	return nil
}

type DbConfig interface {
	ConnectDB() (*gorm.DB, error)
}

type MysqlDBConfig struct {
	Address   string `json:"address"`
	User      string `json:"user"`
	Database  string `json:"database"`
	Password  string `json:"password"`
	ExtParams string `json:"ext_params"`
}

func (c *MysqlDBConfig) ConnectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Password, c.Address, c.Database)
	if len(c.ExtParams) > 0 {
		dsn = dsn + "?" + c.ExtParams
	}
	log.Infof(">>>>>>> Connect Mysql dsn: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		//Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Errorf("Connect to Mysql [%s] failed, config:[%v]", c.Database, c)
		return nil, errors.New(fmt.Sprintf("Connect to Mysql [%s] failed, config:[%v]", c.Database, c))
	}
	log.Infof(">>>>>>> Connect Mysql db [%s] succeed!", c.Database)
	return db, nil
}
