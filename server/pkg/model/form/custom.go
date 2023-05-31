package form

type PageResult struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}
