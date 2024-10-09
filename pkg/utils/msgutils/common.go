package msgutils

type Data map[string]interface{}

type Msg struct {
	Data Data
}

func NewMessage() Msg {
	return Msg{
		Data: make(Data),
	}
}

func (m Msg) Set(key string, value interface{}) Msg {
	m.Data[key] = value
	return m
}

func (m Msg) Done() Data {
	return m.Data
}

func RequestBodyParseErrorResponseMsg() Data {
	return NewMessage().Set("message", "failed to parse request body").Done()
}

func JwtCreateErrorMsg() Data {
	return NewMessage().Set("message", "failed to create jwt token").Done()
}

func SomethingWentWrongMsg() Data {
	return NewMessage().Set("message", "something went wrong").Done()
}

func ExpectationFailedMsg() Data {
	return NewMessage().Set("message", "expectation failed").Done()
}

func AccessForbiddenMsg() Data {
	return NewMessage().Set("message", "access forbidden").Done()
}

func SuccessResponse(data interface{}) Data {
	return NewMessage().Set("message", data).Done()
}

func InvalidDataRequestMsg() Data {
	return NewMessage().Set("message", "invalid data request").Done()
}

func ValidationErrorMsg(err error) Data {
	return NewMessage().Set("message", err.Error()).Done()
}

func ErrorMsg(err error) Data {
	return NewMessage().Set("message", err.Error()).Done()
}
