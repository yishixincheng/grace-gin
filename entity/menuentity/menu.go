package menuentity

import "yishixincheng/grace-gin/model/mdm"

type MenuEntity struct {
	mdm.Menu
	mdm.Dts
	ContentJson string `json:"content_json"`
}