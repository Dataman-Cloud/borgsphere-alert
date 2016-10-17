package dao

import (
	"github.com/Dataman-Cloud/borgsphere-alert/src/model"
)

type Store interface {
	CreateStrategy(strategy *model.Strategy) error
	Test()
}
