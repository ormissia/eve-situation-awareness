package utils

const (
	SystemName   = "eve-situation-awareness"
	DTTimeFormat = "2006-01-02 15"
)

const (
	Hour  = "hour"
	Day   = "day"
	Month = "month"
	Year  = "year"
)

var (
	SQLGroupFormatMap = map[string]string{
		Hour:  "%Y-%m-%d %k",
		Day:   "%Y-%m-%d",
		Month: "%Y-%m",
		Year:  "%Y",
	}
	ResultDTFormatMap = map[string]string{
		Hour:  "2006-01-02 15",
		Day:   "2006-01-02",
		Month: "2006-01",
		Year:  "2006",
	}
)
