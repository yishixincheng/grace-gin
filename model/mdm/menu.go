
package mdm


type Menu struct {
	 // 菜单主键 
	 MenuId	int32	`gorm:"type:int(11); not null" json:"menu_id"`
	 // 菜单名称 
	 MenuName	string	`gorm:"type:varchar(255); not null" json:"menu_name"`
	 // 数据源id 
	 DtsId	string	`gorm:"type:varchar(255); not null" json:"dts_id"`
	 // 数据库名 
	 Db	string	`gorm:"type:varchar(255); not null" json:"db"`
	 // sql语句模板 
	 TplSql	string	`gorm:"type:varchar(255); not null; default:''" json:"tpl_sql"`
}

func (model Menu) TableName() string {
	return "menu"
}

