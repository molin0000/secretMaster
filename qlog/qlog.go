package qlog

import "fmt"

type Priority int32

const (
	Debug       Priority = 0
	Info                 = 10
	InfoSuccess          = 11
	InfoRecv             = 12
	InfoSend             = 13
	Warning              = 20
	Error                = 30
	Fatal                = 40
)

var _addLog func(p int32, logType, reason string) int32

func HandleLog(addLog func(p int32, logType, reason string) int32) {
	_addLog = addLog
}

func AddLog(p Priority, logType, reason string) int32 {
	fmt.Println(reason)
	if _addLog == nil {
		return 0
	}
	return _addLog(int32(p), logType, reason)
}

func Println(msg ...interface{}) {
	formatStr := ""
	for i := 0; i < len(msg); i++ {
		formatStr += "%+v "
	}

	formatStr += "\n"
	AddLog(Debug, "工作日志", fmt.Sprintf(formatStr, msg...))
}
