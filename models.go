package models

type QueryRequest struct {
	Query string `json:"query"`
}

type QueryResult struct {
	Result []map[string]interface{} `json:"result"`
}
