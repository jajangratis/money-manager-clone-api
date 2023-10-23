package web

type MstAccountCreateRequest struct {
	AccountId   string `validate:"required,max=200,min=1" json:"account_id"`
	AccountName string `validate:"required,max=200,min=1" json:"account_name"`
}
