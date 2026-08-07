package core

type LumaError struct {
	IsLumaError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewLumaError(code string, msg string, ctx *Context) *LumaError {
	return &LumaError{
		IsLumaError: true,
		Sdk:              "Luma",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *LumaError) Error() string {
	return e.Msg
}
