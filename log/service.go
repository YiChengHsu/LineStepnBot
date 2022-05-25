package log

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

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
	time.Sleep(2 * time.Second)
	os.Exit(1)
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

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			requestId := uuid.NewString()
			c.Set("requestId", requestId)
			logLabel := fmt.Sprintf("%s%s ", LabelMonitor, requestId)
			Info(logLabel, fmt.Sprintf("[%s] %s %s %s", req.Method, req.URL, req.UserAgent(), req.Referer()))

			if err = next(c); err != nil {
				return
			}

			res := c.Response()
			logContent := fmt.Sprintf(
				"Status=%d ClientAddr=%s Response=%+v", res.Status, c.RealIP(), c.Get("response"),
			)

			if res.Status < 400 {
				Info(logLabel, logContent)
			}

			return
		}
	}
}
