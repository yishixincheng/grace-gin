
package mdm

import (
	"github.com/jinzhu/gorm"
)


type RelMenuRoleField struct {
     gorm.Model
	 // id 
	 Id	int32	`gorm:"type:int(11); not null" json:"id"`
	 // rel_menu_role_id 
	 RelMenuRoleId	int32	`gorm:"type:int(11); not null" json:"rel_menu_role_id"`
	 // 字段（a.store_name或者name之类的用作menu的column的填充） 
	 Field	string	`gorm:"type:varchar(255); not null" json:"field"`
	 // 字段名（中文，展示用） 
	 Name	string	`gorm:"type:varchar(255); not null" json:"name"`
	 // 类型（1、string；2、number） 
	 Type	int8	`gorm:"type:tinyint(4); not null" json:"type"`
	 // 是否开放搜索 
	 CanSearch	int8	`gorm:"type:tinyint(4); not null" json:"can_search"`
	 // 字段最大长度 
	 MaxLength	int32	`gorm:"type:int(11); not null" json:"max_length"`
	 // 字段最小长度 
	 MinLength	int32	`gorm:"type:int(11); not null" json:"min_length"`
}

func (model RelMenuRoleField) TableName() string {
	return "rel_menu_role_field"
}

