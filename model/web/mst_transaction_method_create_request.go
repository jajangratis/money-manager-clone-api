package web

type MstTransactionMethodCreateRequest struct {
	MethodId   string `validate:"required,max=200,min=1" json:"method_id"`
	MethodName string `validate:"required,max=200,min=1" json:"method_name"`
}
