package user

type DeleteReq struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

type DeleteResp struct {
	Code int64           `json:"code"`
	Msg  string          `json:"msg"`
	Data *DeleteRespData `json:"data"`
}
type DeleteRespData struct {
	Result bool `json:"result"`
}

func NewDeleteResp() *DeleteResp {
	return &DeleteResp{
		Data: &DeleteRespData{},
	}
}
