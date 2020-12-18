package Unified_exception_handling_demo

//实现一个异常接口
type UserError string

//实现UserError接口里面的方法
func (e UserError) Error() string {
	return e.Message()
}

//实现UserError接口里面的方法
func (e UserError) Message() string {
	return string(e)
}