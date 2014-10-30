package common

type ErrorInfo struct {
	Message string `json:message`
	Code    int    `json:code`
}
type RetData struct {
	Sendno string `json:sendno`
	MsgID  string `json:msg_id`
	Error  ErrorInfo
}
