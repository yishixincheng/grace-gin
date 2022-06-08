
package mdm

import (
	"github.com/jinzhu/gorm"
	"time"
)


type Role struct {
     gorm.Model
	 // 角色主键 
	 RoleId	int32	`gorm:"type:int(11); not null" json:"role_id"`
	 // 角色名称 
	 Name	string	`gorm:"type:varchar(255); not null" json:"name"`
	 // 创建时间 
	 CreateTs	time.Time	`gorm:"type:timestamp; not null; default:'CURRENT_TIMESTAMP'" json:"create_ts"`
	 // 是否被删除（0、未删除；1、删除） 
	 IsDelete	int8	`gorm:"type:tinyint(4); not null" json:"is_delete"`
}

func (model Role) TableName() string {
	return "role"
}

