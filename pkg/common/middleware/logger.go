package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	c "github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	l := &logrus.Logger{
		Out: os.Stdout, Formatter: &CustomJSONFormatter{},
		Hooks: make(logrus.LevelHooks), Level: logrus.InfoLevel,
	}
	log.Println("Logger OK")
	return &Logger{l}
}

// CustomJSONFormatter is a custom formatter for logrus
type CustomJSONFormatter struct{}

// Format leaves the messages like this:
//
//	 2023-11-02 14:47:44 -> GetUser: user.Deleted -> error, user already deleted
//
//		 {
//		     "path": "/v1/users/8",
//		     "status": 404
//		 }
//
// -
func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	formattedTime := entry.Time.Format("2006-01-02 15:04:05")

	// We only use the custom formatter for errors. TODO rework this!
	if entry.Level > logrus.ErrorLevel {
		return []byte(fmt.Sprintf("%s %s", formattedTime, entry.Message)), nil
	}

	// Entry to JSON
	logJSON, err := json.MarshalIndent(entry.Data, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling fields to JSON -> %v", err)
	}

	// Colorize status code
	isError := false
	status, ok := entry.Data["status"].(int)
	if ok {
		isError = status >= 500
	}

	colorizedJSON := colorizeJSON(string(logJSON), isError)

	// Layout: \n + WhiteBold + [time] + Reset + Red + [message] + \n + White + [JSON] + \n
	layout := "\n%s%s ->%s %s%s\n%s%s\n"
	return []byte(fmt.Sprintf(layout, c.WhiteBold, formattedTime, c.Reset, c.Red, entry.Message, c.White, colorizedJSON)), nil
}

// colorizeJSON just colors the status code in the JSON
func colorizeJSON(json string, isError bool) string {
	beforeStatus := c.RedBold
	if !isError {
		beforeStatus = c.YellowBold
	}
	afterStatus := c.Reset

	// Regular expression to find the status code
	re := regexp.MustCompile(`"status": (\d+)`)

	// Replace the status code with the new string using a regular expression
	return re.ReplaceAllString(json, `"status": `+beforeStatus+`$1`+afterStatus)
}

//

func (l *Logger) Error(msg string, context map[string]interface{}) {
	l.Logger.Error(msg, context)
}
func (l *Logger) Warn(msg string, context map[string]interface{}) {
	l.Logger.Warn(msg, context)
}
func (l *Logger) Info(msg string, context map[string]interface{}) {
	l.Logger.Info(msg, context)
}
func (l *Logger) Debug(msg string, context map[string]interface{}) {
	l.Logger.Debug(msg, context)
}
func (l *Logger) DebugEnabled() bool {
	return l.IsLevelEnabled(logrus.DebugLevel)
}
