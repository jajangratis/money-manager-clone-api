package web

type MstTransactionMethodResponse struct {
	Id          string `json:"id"`
	MethodId    string `json:"method_id"`
	MethodName  string `json:"method_name"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
}
