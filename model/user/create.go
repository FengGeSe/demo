package user

type CreateReq struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}
type CreateResp struct {
	Code int64           `json:"code"`
	Msg  string          `json:"msg"`
	Data *CreateRespData `json:"data"`
}
type CreateRespData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func NewCreateResp() *CreateResp {
	return &CreateResp{
		Data: &CreateRespData{},
	}
}
