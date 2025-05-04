package utility

import "time"

func FormatDateForDynamoDB(d time.Time) string {
	return d.Format(time.RFC3339)
}
