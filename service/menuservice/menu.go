package menuservice

import (
	"encoding/json"
	"fmt"
	"yishixincheng/grace-gin/dao/dtsdao"
	"yishixincheng/grace-gin/dao/menudao"
	"yishixincheng/grace-gin/entity/menuentity"
	"yishixincheng/grace-gin/pkg/dbdriver"
	"yishixincheng/grace-gin/vo/input/menuvo"
	"strconv"
)

type MenuService struct {
	menuDao menudao.IMenuDao
	dtsDao dtsdao.IDtsDao
}

func NewMenuService() *MenuService {
	return &MenuService{
		menuDao: menudao.NewMenuDao(""),
		dtsDao: dtsdao.NewDtsDao(""),
	}
}

func (s *MenuService) GetMenuList(request menuvo.IMenuListRequest) []menuentity.MenuEntity {
	list := s.menuDao.GetMenuList(request.Keyword, request.Page, request.PageSize)
	var menuList []menuentity.MenuEntity

	if len(list) > 0 {
		for i := range list {
			dtsId := list[i].DtsId
			menu := menuentity.MenuEntity{}
			menu.Menu = list[i]
			if dtsId != "" {
				id, err := strconv.ParseInt(dtsId, 10, 32)
				if err != nil {
					continue
				}
				dts := s.dtsDao.QueryById(int32(id))
				menu.Dts = dts
				// menu.ContentJson = s.getContentJson(menu)
			}
			menuList = append(menuList, menu)
		}
	}

	fmt.Println(menuList)

	return menuList

}

func (s *MenuService) getContentJson(menu menuentity.MenuEntity) string {
	if menu.TplSql == "" {
		return ""
	}
	if menu.Dts.DtsId == 0 {
		return ""
	}
	// 创建链接执行语句
	config := dbdriver.ConnectConfig{
		Name: menu.Dts.Ip,
		Host: menu.Dts.Ip,
		Port: menu.Dts.Port,
		Driver: "mysql",
		Database: menu.Db,
		Password: menu.Password,
		Username: menu.User,
		Charset: "utf8mb4",
		Loc: "Asia/Shanghai",
	}
	db := dbdriver.OpenConnect(config)
	var result []map[string]interface{}
	db.Raw(menu.TplSql).Scan(&result)

	marshal, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	return string(marshal)
}