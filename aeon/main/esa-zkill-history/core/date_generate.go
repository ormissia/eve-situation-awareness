package core

func CurrentDay() func() int {
	// ZKill有数据的第一天
	firstDay := 20071205
	// firstDay := 20211205
	return func() int {
		currentDay := firstDay
		firstDay++
		return currentDay
	}
}
