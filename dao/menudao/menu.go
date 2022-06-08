package menudao

import (
	"github.com/jinzhu/gorm"
	"yishixincheng/grace-gin/model/mdm"
	"yishixincheng/grace-gin/pkg/dbdriver"
)

type IMenuDao interface {
	GetMenuList(keyword string, page uint, pageSize uint) []mdm.Menu
}


type MenuDao struct {
	DB *gorm.DB
}

func (d *MenuDao) GetMenuList(keyword string, page uint, pageSize uint) []mdm.Menu {

	var menus []mdm.Menu

	tx := d.DB.Debug().Offset((page - 1) * pageSize).Limit(pageSize)

	if keyword != "" {
		tx = tx.Where("menu_name like ?", "%" +keyword + "%")
	}

	tx.Find(&menus)

	return menus
}

func NewMenuDao(scheme string) IMenuDao {
	return &MenuDao{
		DB: dbdriver.GetConnect(scheme),
	}
}