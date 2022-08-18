package response

type Response struct {
	Code int
	Msg string
	Data interface{}
}
func FailResponse(Code int,Msg string) Response {
	return Response{
		Code: Code,
		Msg:  Msg,
	}
}
func SuccessResponse(Code int,Msg string,data interface{}) Response {
	return Response{
		Code: Code,
		Msg:  Msg,
		Data: data,
	}
}


