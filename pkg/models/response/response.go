package response

type Response struct{
	OriginalURL string `json:"originalURL"`
	ShortenedURL string `json:"shortenedURL"`
}