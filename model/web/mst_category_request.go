package web

type CarrierScoreRequest struct {
	Id           string `json:"id"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}
