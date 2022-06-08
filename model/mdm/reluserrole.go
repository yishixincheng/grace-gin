
package mdm

import (
	"github.com/jinzhu/gorm"
)


type RelUserRole struct {
     gorm.Model
	 // id 
	 Id	int32	`gorm:"type:int(11); not null" json:"id"`
	 // user_id 
	 UserId	int32	`gorm:"type:int(11); not null" json:"user_id"`
	 // role_id 
	 RoleId	int32	`gorm:"type:int(11); not null" json:"role_id"`
}

func (model RelUserRole) TableName() string {
	return "rel_user_role"
}

