package timeutil

import "time"

func TimeLocal() {
	sh, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		sh = time.FixedZone("CST", 8*3600)
	}
	time.Local = sh
}

//距离某个时间戳的天数
func DaysSinceTimestamp(timestamp int64) int64 {
	// 当前时间戳
	now := time.Now().Unix()
	// 计算时间戳相差的秒数
	secondsDiff := now - timestamp
	// 将秒数转换为天数
	return secondsDiff / (3600 * 24)
}

// 获取某天的零点时间戳（秒级）
func GetZeroTimestamp(year int, month int, day int) int64 {
	t := time.Date(year, GetMonth(month), day, 0, 0, 0, 0, time.Local)
	return t.Unix()
}

// 获取当日零点时间戳（秒级）
func GetTodayZeroTimestamp() int64 {
	// 获取当前时间
	now := time.Now()
	// 获取当日零点时间
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	// 将零点时间转换为时间戳
	return startOfDay.Unix()
}

// 获取某个时间戳所在当天的零点时间戳
func GetStartOfDayTimestamp(timestamp int64) int64 {
	t := time.Unix(timestamp, 0)
	year, month, day := t.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return startOfDay.Unix()
}

// 获取前N天的零点时间戳（秒级）
func GetZeroTimestampPreDays(n int) int64 {
	// 获取当前时间
	now := time.Now()
	// 获取前n天的同一时间
	nDaysAgo := now.AddDate(0, 0, -n)
	// 将时间设置为当天的零点
	zeroHour := time.Date(nDaysAgo.Year(), nDaysAgo.Month(), nDaysAgo.Day(), 0, 0, 0, 0, nDaysAgo.Location())
	// 返回零点时间戳
	return zeroHour.Unix()
}

// 今天是周几
func GetTodayWeekday() int {
	// 获取当前的时间
	now := time.Now()

	// 获取今天是周几
	weekday := now.Weekday()

	switch weekday {
	case time.Sunday:
		return 7
	case time.Monday:
		return 1
	case time.Tuesday:
		return 2
	case time.Wednesday:
		return 3
	case time.Thursday:
		return 4
	case time.Friday:
		return 5
	case time.Saturday:
		return 6
	default:
		return 0
	}
}

func GetMonth(mon int) time.Month {
	switch mon {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	case 12:
		return time.December
	default:
		return time.January
	}
}
