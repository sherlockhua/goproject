package model

type Long2ShortRequest struct {
	OriginUrl string `json:"origin_url"`
}


type Long2ShortResponse struct {
	ShortUrl string `json:"short_url"`
}