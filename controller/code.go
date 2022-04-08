package controller

type Statuscode int

const (
	CodeSuccess Statuscode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNoExist
	CodeInvalidPassword
	CodeBusyServer
	CodeAuthFail
)

var codeMsgMap = map[Statuscode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "invalid parameters",
	CodeUserExist:       "the user has already existed",
	CodeUserNoExist:     "No such user",
	CodeInvalidPassword: "invalid user or password",
	CodeBusyServer:      "the server is now busy",
	CodeAuthFail:        "fail to authorize",
}

func (s Statuscode) Msg() string {
	msg, ok := codeMsgMap[s]
	if ok {
		return msg
	} else {
		return codeMsgMap[CodeBusyServer]
	}
}
