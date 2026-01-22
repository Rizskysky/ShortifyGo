package entity

type Url struct {
	Code      string `json:"code"`
	OriginUrl string `json:"origin_url"`
	Clicks    int    `json:"clicks"`
}

type CreateUrlRequest struct {
	Url string `json:"url"`
}

type CreateUrlResponse struct {
	Url  string `json:"url"`
	Code string `json:"code"`
}
