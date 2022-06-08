
package mdm

import (
	"github.com/jinzhu/gorm"
	"time"
)


type Admin struct {
     gorm.Model
	 // id 
	 Id	int32	`gorm:"type:int(11); not null" json:"id"`
	 // 账号 
	 Account	string	`gorm:"type:varchar(255); not null" json:"account"`
	 // 密码 
	 Passwd	string	`gorm:"type:varchar(255); not null" json:"passwd"`
	 // 头像 
	 Avatar	string	`gorm:"type:varchar(255); not null" json:"avatar"`
	 // dd_userid 
	 DdUserid	string	`gorm:"type:varchar(255); not null" json:"dd_userid"`
	 // name 
	 Name	string	`gorm:"type:varchar(255); not null" json:"name"`
	 // phone 
	 Phone	string	`gorm:"type:varchar(255); not null" json:"phone"`
	 // 1 普通  2组长  3客服 99超级管理员 
	 RoleId	int8	`gorm:"type:tinyint(4); not null; default:'1'" json:"role_id"`
	 // 1 正常 2删除 
	 Status	int8	`gorm:"type:tinyint(4); not null; default:'1'" json:"status"`
	 // create_time 
	 CreateTime	time.Time	`gorm:"type:timestamp; not null; default:'CURRENT_TIMESTAMP'" json:"create_time"`
}

func (model Admin) TableName() string {
	return "admin"
}

