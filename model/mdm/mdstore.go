
package mdm

import (
	"github.com/jinzhu/gorm"
)


type MdStore struct {
     gorm.Model
	 // id 
	 Id	int32	`gorm:"type:int(11); not null" json:"id"`
	 // 门店编码（主数据-来源于加盟商管理系统编码） 
	 MdStoreNo	string	`gorm:"type:varchar(64); not null" json:"md_store_no"`
	 // 门店名称（主数据-来源于加盟商管理系统） 
	 MdStoreName	string	`gorm:"type:varchar(255); not null" json:"md_store_name"`
}

func (model MdStore) TableName() string {
	return "md_store"
}

