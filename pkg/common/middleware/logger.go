package middleware

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

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

	if entry.Message == "" {
		return nil, nil
	}

	// We always output first the time
	formattedTime := entry.Time.Format("2006-01-02 15:04:05")

	// If it's a New Relic log, we put it in BlueBold
	if isNewRelicLog(entry.Message) {
		return []byte(fmt.Sprintf("%s%s %s[New Relic]%s %s%s", c.WhiteBold, formattedTime, c.BlueBold, c.Reset, c.White, entry.Message)), nil
	}

	// If it's a Prometheus log, we put it in MagentaBold
	if isPrometheusLog(entry.Data) {
		return []byte(fmt.Sprintf("%s%s %s[Prometheus]%s %s%s", c.WhiteBold, formattedTime, c.MagentaBold, c.Reset, c.White, entry.Message)), nil
	}

	if isGinLog(entry.Message) {
		message, _ := strings.CutPrefix(entry.Message, "[GIN-debug]")
		return []byte(fmt.Sprintf("%s%s %s[GIN]%s %s%s\n", c.WhiteBold, formattedTime, c.GreenBold, c.Reset, c.White, message)), nil
	}

	// We only use the custom formatter for errors. TODO rework this!
	if entry.Level > logrus.ErrorLevel {
		return []byte(fmt.Sprintf("%s%s%s %s%s%s", c.WhiteBold, formattedTime, c.Reset, c.White, entry.Message, c.Reset)), nil
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
	log := l.Logger.WithField("msg", msg)
	for k, v := range context {
		log = log.WithField(k, v)
	}

	if !strings.Contains(msg, "\n") {
		msg += "\n"
	}

	log.Info(msg)
}
func (l *Logger) Debug(msg string, context map[string]interface{}) {
	log := l.Logger.WithField("msg", msg)
	for k, v := range context {
		log = log.WithField(k, v)
	}

	if !strings.Contains(msg, "\n") {
		msg += "\n"
	}

	log.Debug(msg)
}
func (l *Logger) DebugEnabled() bool {
	return l.IsLevelEnabled(logrus.DebugLevel)
}

func isNewRelicLog(msg string) bool {
	return msg == "application created\n" || msg == "final configuration\n" || msg == "application connected\n" || msg == "collector message\n"
}

func isPrometheusLog(data logrus.Fields) bool {
	val, ok := data["from"]
	return ok && val == "Prometheus"
}

func isGinLog(msg string) bool {
	return strings.Contains(msg, "[GIN-debug]") || strings.Contains(msg, "GIN_MODE") || strings.Contains(msg, "gin.SetMode")
}
