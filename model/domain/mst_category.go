package domain

type MstCategory struct {
	Id           string `json:"id"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreatedDate  string `json:"created_date"`
	UpdatedDate  string `json:"updated_date"`
}

type MstCategoryResultDb struct {
	Id           string `json:"id"`
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	CreatedDate  string `json:"created_date"`
	UpdatedDate  string `json:"updated_date"`
}
