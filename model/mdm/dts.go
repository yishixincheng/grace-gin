
package mdm

import "time"

type Dts struct {
	 // dts_id 
	 DtsId	int32	`gorm:"type:int(11); not null" json:"dts_id"`
	 // 数据源名称 
	 Name	string	`gorm:"type:varchar(255); not null" json:"name"`
	 // 类型（MYSQL） 
	 Type	string	`gorm:"type:varchar(255); not null" json:"type"`
	 // 数据库连接（ip） 
	 Ip	string	`gorm:"type:varchar(255); not null" json:"ip"`
	 // 端口 
	 Port	string	`gorm:"type:varchar(255); not null" json:"port"`
	 // 用户名 
	 User	string	`gorm:"type:varchar(255); not null" json:"user"`
	 // 密码 
	 Password	string	`gorm:"type:varchar(255); not null" json:"password"`
	 // 创建时间 
	 CreateTs	time.Time	`gorm:"type:timestamp; not null; default:'CURRENT_TIMESTAMP'" json:"create_ts"`
}

func (model Dts) TableName() string {
	return "dts"
}

