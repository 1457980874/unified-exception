package Error_Mould

//自定义一个异常
//定义一个异常接口
type UserError interface {
	error
	Message() string
}
