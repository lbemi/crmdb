package types

type Deployment struct {
}

type PageQuery struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}
