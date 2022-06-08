
package mdm

import (
	"github.com/jinzhu/gorm"
)


type User struct {
     gorm.Model
	 // 钉钉userId 
	 UserId	int32	`gorm:"type:int(11); not null" json:"user_id"`
	 // 钉钉授权手机号 
	 Mobile	string	`gorm:"type:varchar(11); not null" json:"mobile"`
	 // 行限制条件（1、无限制；2、角色限制；） 
	 YLimit	int8	`gorm:"type:tinyint(4); not null" json:"y_limit"`
	 // 是否被删除（0、未删除；1、已删除） 
	 IsDelete	int8	`gorm:"type:tinyint(4); not null" json:"is_delete"`
}

func (model User) TableName() string {
	return "user"
}

