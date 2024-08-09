package dto

type PageRequest struct {
	Page     int    `json:"page" form:"page,default=1"`
	PageSize int    `json:"page_size" form:"page_size,default=10"`
	Keyword  string `json:"keyword" form:"keyword"`
}

type PageResponse struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}
