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
// 		- 2023-11-02 14:47:44 [•] Middlewares OK
//		- 2023-11-02 14:47:44 [New Relic] application created
// 		- 2023-11-02 14:47:44 [SQL] SELECT DATABASE() (-1 rows affected) [0.280ms]
//
//------------------------------------------------------*/

type CustomFormatter struct{}

var (
	timeFormat = "2006-01-02 15:04:05"

	entryKeyRows = "rows"
	entryKeyFrom = "from"

	defaultSourceAppendage = fmt.Sprintf(" %s[•]", DefaultColor)
	sqlSourceAppendage     = fmt.Sprintf(" %s[SQL]", GormColor)

	ginPrefix = "[GIN-debug]"
)

/*--------------------
//   Main Formatter
//-------------------*/

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// Discard empty messages
	if entry.Message == "" || entry.Message == "\n" {
		return nil, nil
	}

	// Log from Sources: New Relic, Prometheus, Gin, Gorm(SQL)
	if ok, log := formatLogFromSources(entry, entry.Time.Format(timeFormat)); ok {
		return log, nil
	}

	// Default info formatter, pretty standard.
	if entry.Level > logrus.ErrorLevel {
		return formatLogDefault(entry.Message, entry.Time.Format(timeFormat)), nil
	}

	if entry.Data["status"] != nil && entry.Data["path"] != nil && entry.Data["method"] != nil {
		// Only our custom errors should get up here.
		// We parse the JSON and colorize the status code.
		return formatLogCustomError(entry.Message, parseAndColorJSON(entry.Data), entry.Time.Format(timeFormat)), nil
	}

	return formatLogDefaultError(entry.Message, entry.Time.Format(timeFormat)), nil
}

/*--------------------
//   Custom Errors
//-------------------*/

func formatLogCustomError(message, json, time string) []byte {
	customErrorLayout := "\n%s%s ->%s %sError %s%s%s\n%s%s%s\n"
	return []byte(fmt.Sprintf(customErrorLayout, WhiteBold, time, Nil, RedBold, Nil, Red, message, White, json, Nil))
}

/*--------------------
//      Default
//-------------------*/

func formatLogDefault(message, time string) []byte {
	defaultLayout := "%s%s %s[•] %s%s%s%s"
	return []byte(fmt.Sprintf(defaultLayout, WhiteBold, time, DefaultColor, Nil, White, message, Nil))
}

func formatLogDefaultError(message, time string) []byte {
	defaultLayout := "%s%s %s[Error] %s%s%s%s"
	return []byte(fmt.Sprintf(defaultLayout, WhiteBold, time, RedBold, Nil, White, message, Nil))
}

/*--------------------
//   From Sources
//-------------------*/

func formatLogFromSources(entry *logrus.Entry, time string) (bool, []byte) {

	if isFromSource(entry.Data, common.NewRelic) {
		return true, formatLogFromSource(common.NewRelic.StrNice(), entry.Message, NewRelicColor, time)
	}

	if isFromSource(entry.Data, common.Prometheus) {
		return true, formatLogFromSource(common.Prometheus.StrNice(), entry.Message, PrometheusColor, time)
	}

	if isFromSource(entry.Data, common.Gin) {
		message, _ := strings.CutPrefix(entry.Message, ginPrefix)
		return true, formatLogFromSource(common.Gin.StrNice(), message, GinColor, time)
	}

	if isFromSource(entry.Data, common.Gorm) {
		if _, isSQL := entry.Data[entryKeyRows]; isSQL {
			return true, formatLogSQL(entry.Message, time, entry.Data)
		}
		return true, formatLogFromSource(common.Gorm.StrNice(), entry.Message, GormColor, time)
	}

	return false, []byte{}
}

func formatLogFromSource(source, message, color, time string) []byte {
	fromSourceLayout := "%s%s %s[%s]%s %s%s%s"
	return []byte(fmt.Sprintf(fromSourceLayout, WhiteBold, time, color, source, Nil, White, message, Nil))
}

/*--------------------
//        SQL
//-------------------*/

func formatLogSQL(message, time string, fields logrus.Fields) []byte {

	prefix := fmt.Sprintf("\n%s%s ", WhiteBold, time) // Log beginning
	message = strings.ReplaceAll(message, "\n", "")
	rows, elapsed, err := fields["rows"], fields["elapsed"], fields["err"]

	// Print Error or SQL
	if err != nil {
		sqlLayout := prefix + "%s[SQL Error]%s %s%s%s (%d rows affected) %s[%.3fms] %s%v%s\n"
		return []byte(fmt.Sprintf(sqlLayout, RedBold, Nil, White, message, Gray, rows, Green, elapsed, Red, err, Nil))
	}

	sqlLayout := prefix + "%s[SQL]%s %s%s%s (%d rows affected) %s[%.3fms]%s\n"
	return []byte(fmt.Sprintf(sqlLayout, GormColor, Nil, White, message, Gray, rows, Green, elapsed, Nil))
}

/*------------------------
//    Sources Helpers
//---------------------*/

func isFromSource(data logrus.Fields, source common.LogSource) bool {
	val, ok := data[entryKeyFrom]
	return ok && val == source.Str()
}

/*----------------------
//       Colors
//--------------------*/

const (
	Nil     = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"

	RedBold     = "\033[31;1m"
	GreenBold   = "\033[32;1m"
	YellowBold  = "\033[33;1m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	WhiteBold   = "\033[37;1m"
	CyanBold    = "\033[36;1m"

	DefaultColor    = CyanBold
	GinColor        = GreenBold
	GormColor       = YellowBold
	NewRelicColor   = BlueBold
	PrometheusColor = MagentaBold
)

// parseAndColorJSON takes a logrus.Fields and returns it as JSON with the status code colored accordingly
func parseAndColorJSON(fields logrus.Fields) string {
	json, err := json.MarshalIndent(fields, "", "    ")
	if err != nil {
		return fmt.Sprintf("error marshalling fields to JSON -> %v", err)
	}

	status, ok := fields["status"].(int)
	yellowOrRed := YellowBold
	if ok && status >= 500 {
		yellowOrRed = RedBold
	}

	return regexp.MustCompile(`"status": (\d+)`).ReplaceAllString(string(json), `"status": `+yellowOrRed+`$1`+Nil)
}
