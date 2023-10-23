package web

type MstAccountTypeRequest struct {
	Id          string `json:"id"`
	AccountId   string `json:"account_id"`
	AccountName string `json:"account_name"`
}
