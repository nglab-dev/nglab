package dto

type PaginationParam struct {
	Page     int    `json:"page" form:"page,default=1"`
	PageSize int    `json:"page_size" form:"page_size,default=10"`
	Keyword  string `json:"keyword" form:"keyword"`
}

type PaginationResult struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	Items    interface{} `json:"items"`
}
