package models

type URLCheckRequest struct {
	Links []string `json:"links"`
}

type URLCheckResponse struct {
	ID      int               `json:"id"`
	Results map[string]string `json:"results"`
}

type URLRecord struct {
	ID      int               `json:"id"`
	Links   []string          `json:"links"`
	Results map[string]string `json:"results"`
}

type ReportRequest struct {
	LinksNum []int `json:"links_num"`
}

type ReportData struct {
	ID     int
	Link   string
	Status string
}
