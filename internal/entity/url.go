package entity

type URL struct {
	ID        string `json:"id"`       // Short Code (contoh: "xYz12")
	Original  string `json:"original"` // Link Asli (contoh: "https://shopee.co.id")
}