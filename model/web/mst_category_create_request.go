package web

type MstCategoryCreateRequest struct {
	CategoryId   string `validate:"required,max=200,min=1" json:"category_id"`
	CategoryName string `validate:"required,max=200,min=1" json:"category_name"`
}
