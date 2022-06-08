package menuvo

type IMenuListRequest struct {
	Page uint `json:"page" binding:"required,min=1"`
	PageSize uint `json:"page_size" binding:"required,max=100"`
	Keyword string `json:"keyword"`
}