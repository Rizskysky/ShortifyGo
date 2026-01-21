package entity

type URL struct {
	//ID menyimpan shorted url + sebagai identifer
	ID string `json:"id"`
	//Original url yang di shorted
	OriginalURL string `json:"original_url"`
}

type CreateShortUrlRequest struct {
	URL string `json:"url"`
}

type CreateShortUrlResponse struct {
	Code string `json:"code"`
}

type BadResponse struct {
	Message string `json:"message"`
}
