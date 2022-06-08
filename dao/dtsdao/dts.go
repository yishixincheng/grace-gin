package dtsdao

import (
	"github.com/jinzhu/gorm"
	"yishixincheng/grace-gin/model/mdm"
	"yishixincheng/grace-gin/pkg/dbdriver"
)

type IDtsDao interface {
	QueryById(id int32) mdm.Dts
}

type DtsDao struct {
	DB *gorm.DB
}

func (d *DtsDao) QueryById(id int32) mdm.Dts {
	var dts mdm.Dts
	d.DB.Debug().Where("dts_id = ?", id).Find(&dts)

	return dts
}

func NewDtsDao(scheme string) IDtsDao {
	return &DtsDao{
		DB: dbdriver.GetConnect(scheme),
	}
}
