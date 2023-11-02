package middleware

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	c "github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	l := &logrus.Logger{
		Out: os.Stdout, Formatter: &CustomJSONFormatter{},
		Hooks: make(logrus.LevelHooks), Level: logrus.DebugLevel,
	}
	return l
}

// CustomJSONFormatter is a custom formatter for logrus
type CustomJSONFormatter struct{}

// Format leaves the messages like this:
//
//	 2023-11-02 14:47:44 -> GetUser: user.Deleted -> error, user already deleted
//
//		 {
//		     "path": "/v1/users/:user_id",
//		     "status": 404
//		 }
//
// -
func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// Entry to JSON
	logJSON, err := json.MarshalIndent(entry.Data, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("error marshalling fields to JSON -> %v", err)
	}

	// Colorize status code
	isError := entry.Data["status"].(int) >= 500
	colorizedJSON := colorizeJSON(string(logJSON), isError)
	formattedTime := entry.Time.Format("2006-01-02 15:04:05")

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
