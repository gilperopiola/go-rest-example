package middleware

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/sirupsen/logrus"
)

/* --- Our custom error messages log like this:
//
//	 2023-11-02 14:47:44 -> GetUser: user.Deleted -> error, user already deleted
//
//		 {
//		     "path": "/v1/users/8",
//		     "status": 404
//		 }
//
// 	Other types of logs have other formats:
//
// 		- Middlewares OK
//		- [New Relic] application created
// 		- [SQL] SELECT DATABASE() (-1 rows affected) [0.280ms]
//
//------------------------------------------------------*/

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// Discard empty messages
	if entry.Message == "" {
		return nil, nil
	}

	formattedTime := entry.Time.Format("2006-01-02 15:04:05")

	// Sources: New Relic, Prometheus, Gin, Gorm
	if ok, log := getLogFromSources(entry, formattedTime); ok {
		return log, nil
	}

	// If it's less than an error, we log it without much formatting
	if entry.Level > logrus.ErrorLevel {
		return formatLogDefault(entry.Message, formattedTime), nil
	}

	// Only our custom errors should get up to this point.
	// We parse the JSON and colorize the status code.
	// Layout: \n + WhiteBold + [time] + Reset + Red + [message] + \n + White + [JSON] + \n
	return formatLogCustomError(entry.Message, parseAndColorJSON(entry.Data), formattedTime), nil
}

func isFromSource(data logrus.Fields, source common.LogSource) bool {
	val, ok := data["from"]
	return ok && val == source.String()
}

func getLogFromSources(entry *logrus.Entry, time string) (bool, []byte) {

	if isFromSource(entry.Data, common.NewRelic) {
		return true, formatLogFromSource("New Relic", entry.Message, BlueBold, time)
	}

	if isFromSource(entry.Data, common.Prometheus) {
		return true, formatLogFromSource("Prometheus", entry.Message, MagentaBold, time)
	}

	if isFromSource(entry.Data, common.Gin) {
		message, _ := strings.CutPrefix(entry.Message, "[GIN-debug]")
		return true, formatLogFromSource("Gin", message, GreenBold, time)
	}

	if isFromSource(entry.Data, common.Gorm) {
		return true, formatLogSQL(entry.Message, time, entry.Data)
	}

	return false, []byte{}
}

func formatLogCustomError(message, json, time string) []byte {
	layout := "\n%s%s ->%s %s%s\n%s%s%s\n"
	return []byte(fmt.Sprintf(layout, WhiteBold, time, Nil, Red, message, White, json, Nil))
}

func formatLogFromSource(source, message, color, time string) []byte {
	layout := "%s%s %s[%s]%s %s%s%s"
	return []byte(fmt.Sprintf(layout, WhiteBold, time, color, source, Nil, White, message, Nil))
}

func formatLogSQL(message, time string, fields logrus.Fields) []byte {

	prefix := fmt.Sprintf("\n%s%s ", WhiteBold, time) // Log beginning
	message = strings.ReplaceAll(message, "\n", "")
	rows, elapsed, err := fields["rows"], fields["elapsed"], fields["err"]

	// Print Error or SQL
	if err != nil {
		layout := prefix + "%s[SQL Error]%s %s%s%s (%d rows affected) %s[%.3fms] %s%v%s\n"
		return []byte(fmt.Sprintf(layout, RedBold, Nil, White, message, Gray, rows, Green, elapsed, Red, err, Nil))
	}

	layout := prefix + "%s[SQL]%s %s%s%s (%d rows affected) %s[%.3fms]%s\n"
	return []byte(fmt.Sprintf(layout, YellowBold, Nil, White, message, Gray, rows, Green, elapsed, Nil))
}

func formatLogDefault(message, time string) []byte {
	return []byte(fmt.Sprintf("%s%s %s%s%s%s", WhiteBold, time, Nil, White, message, Nil))
}

// parseAndColorJSON takes a logrus.Fields and returns it as JSON with the status code colored accordingly
func parseAndColorJSON(fields logrus.Fields) string {
	json, err := json.MarshalIndent(fields, "", "    ")
	if err != nil {
		return fmt.Sprintf("error marshalling fields to JSON -> %v", err)
	}

	status, ok := fields["status"].(int)
	statusColor := YellowBold
	if ok && status >= 500 {
		statusColor = RedBold
	}

	return regexp.MustCompile(`"status": (\d+)`).ReplaceAllString(string(json), `"status": `+statusColor+`$1`+Nil)
}

// Colors
const (
	Nil         = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	Gray        = "\033[90m"
	RedBold     = "\033[31;1m"
	GreenBold   = "\033[32;1m"
	YellowBold  = "\033[33;1m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	WhiteBold   = "\033[37;1m"
)
