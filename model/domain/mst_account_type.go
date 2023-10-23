package domain

type MstAccountType struct {
	Id          string `json:"id"`
	AccountId   string `json:"account_id"`
	AccountName string `json:"account_name"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}
