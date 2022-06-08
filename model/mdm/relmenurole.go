
package mdm

import (
	"github.com/jinzhu/gorm"
)


type RelMenuRole struct {
     gorm.Model
	 // rel_menu_role_id 
	 RelMenuRoleId	int32	`gorm:"type:int(11); not null" json:"rel_menu_role_id"`
	 // 菜单id 
	 MenuId	int32	`gorm:"type:int(11); not null" json:"menu_id"`
	 // 角色id 
	 RoleId	int32	`gorm:"type:int(11); not null" json:"role_id"`
}

func (model RelMenuRole) TableName() string {
	return "rel_menu_role"
}

