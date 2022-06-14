package log

import (
	"encoding/json"
	"fmt"
	"time"
)

var timezone *time.Location = time.Local

func SetTimezone(loc *time.Location) {
	timezone = loc
}

func Info(args ...interface{}) {
	log(levelInfo, args...)
}

func Error(args ...interface{}) {
	log(levelError, args...)
}

func Fatal(args ...interface{}) {
	log(levelFatal, args...)
}

func Debug(args ...interface{}) {
	data, _ := json.MarshalIndent(args, "", " ")
	log(levelInfo)
	fmt.Println(string(data))
}

func log(level string, args ...interface{}) {
	msg := fmt.Sprintf(
		"%s %s",
		level,
		time.Now().In(timezone).Format("2006-01-02T15:04:05.000Z07:00"),
	)
	for _, arg := range args {
		msg = fmt.Sprintf("%s %+v", msg, arg)
	}

	fmt.Println(msg)
}
