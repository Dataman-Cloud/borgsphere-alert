package mysql

import (
	"github.com/Dataman-Cloud/borgsphere-alert/src/dao"
	"github.com/Dataman-Cloud/borgsphere-alert/src/model"
	"github.com/Dataman-Cloud/borgsphere-alert/src/utils/config"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattes/migrate/driver/mysql"
)

type mysqlStore struct {
	dao.Store
	Client *gorm.DB
}

func InitMysqlStore() *mysqlStore {
	if db, err := gorm.Open(config.GetConfig().AlertDbDriver, config.GetConfig().AlertDbDSN); err == nil {
	} else {
	}
	return &mysqlStore{}
}

func (ms *mysqlStore) CreateStrategy(strategy *model.Strategy) error {
	log.Info("create strategy")
	return nil
}

func (ms *mysqlStore) Test() {
	log.Info("this is a test")
}
