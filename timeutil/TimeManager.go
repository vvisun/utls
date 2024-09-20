package timeutil

import (
	"math"
	"time"
)

// 定时器管理工具, 每天 5 点刷新
var StartTime = time.Date(2023, 1, 1, 5, 0, 0, 0, time.Local).UnixMilli()

var mondayTime = time.Date(2023, 1, 2, 5, 0, 0, 0, time.Local)
var tuesdayTime = time.Date(2023, 1, 3, 5, 0, 0, 0, time.Local)
var wednesdayTime = time.Date(2023, 1, 4, 5, 0, 0, 0, time.Local)
var thursdayTime = time.Date(2023, 1, 5, 5, 0, 0, 0, time.Local)
var fridayTime = time.Date(2023, 1, 6, 5, 0, 0, 0, time.Local)
var saturdayTime = time.Date(2023, 1, 7, 5, 0, 0, 0, time.Local)
var sundayTime = time.Date(2023, 1, 8, 5, 0, 0, 0, time.Local)

const OneDayOfMill = 24 * 3600 * 1000
const OneWeekOfMill = 7 * 24 * 3600 * 1000

type StateVar struct {
	Id         StateVarType
	Val        int64
	SaveType   SaveMode
	CreateTime int64 // 起始时间
}

func GetDayOfNow() int64 {
	now := time.Now().UnixMilli()
	diff := (now - StartTime) / OneDayOfMill
	return diff
}

func GetDayOfWeek(weekDay time.Weekday) int64 {
	nowDay := time.Now()
	if nowDay.Hour() < 5 {
		nowDay = nowDay.AddDate(0, 0, -1)
	}
	now := time.Date(nowDay.Year(), nowDay.Month(), nowDay.Day(), 5, 0, 0, 0, time.Local)
	if now.Weekday() == weekDay {
		return now.UnixMilli()
	} else {
		offsetDays := int(weekDay - now.Weekday())
		absDay := int(math.Abs(float64(weekDay - now.Weekday())))
		if offsetDays > 0 {
			offsetDays = absDay - 7
		}
		days := now.AddDate(0, 0, offsetDays)
		return days.UnixMilli()
	}
}

func GetWeekOfMonday() int64 {
	return getWeekOfNow(time.Monday)
}

func GetWeekOfTuesday() int64 {
	return getWeekOfNow(time.Tuesday)
}

func GetWeekOfWednesday() int64 {
	return getWeekOfNow(time.Tuesday)
}

func GetWeekOfThursday() int64 {
	return getWeekOfNow(time.Tuesday)
}

func GetWeekOfFriday() int64 {
	return getWeekOfNow(time.Tuesday)
}

func GetWeekOfSaturday() int64 {
	return getWeekOfNow(time.Tuesday)
}

func GetWeekOfSunday() int64 {
	return getWeekOfNow(time.Tuesday)
}

func getWeekOfNow(week time.Weekday) int64 {
	now := GetDayOfWeek(week)
	diff := (now - StartTime) / OneWeekOfMill
	return diff
}

func NewStateVar(id StateVarType, val int64, saveMode SaveMode) *StateVar {
	var createTime int64 = 0
	switch saveMode {
	case SaveOneDay:
		{
			createTime = GetDayOfNow()
		}
	case SaveOneWeekOfMonday:
		{
			createTime = GetWeekOfMonday()
		}
	}
	return &StateVar{
		Id:         id,
		Val:        val,
		SaveType:   saveMode,
		CreateTime: createTime,
	}
}

func NowToTomorrow() int64 {
	now := time.Now()
	calcNow := now
	if calcNow.Hour() < 5 {
		tomorrow := time.Date(calcNow.Year(), calcNow.Month(), calcNow.Day(), 5, 0, 0, 0, calcNow.Location())
		return tomorrow.Sub(calcNow).Milliseconds()
	}
	// 获取明天的日期
	tomorrow := calcNow.Add(24 * time.Hour)
	// 设置明天凌晨5点的时间
	tomorrow5AM := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 5, 0, 0, 0, tomorrow.Location())
	return tomorrow5AM.Sub(now).Milliseconds()
}

func IsNextDay(lastDateNum int64) bool {
	if lastDateNum <= 0 {
		return true
	}
	nowDay := time.Now()
	lastDate := time.UnixMilli(lastDateNum)
	if lastDate.Day() == nowDay.Day() && lastDate.Hour() >= 5 && nowDay.Hour() >= 5 { // 同一天且都在5点后
		return false
	}
	nextDay := lastDate.Add(24 * 3600 * time.Second)
	return (nowDay.UnixMilli()-nextDay.UnixMilli()) > 0 || nowDay.Hour() >= 5 // 非同一天或当前时间在5点后
}

func NowToNextWeek(weekday time.Weekday) int64 {
	nowDay := time.Now()
	if nowDay.Hour() < 5 {
		nowDay = nowDay.AddDate(0, 0, -1)
	}
	today := time.Date(nowDay.Year(), nowDay.Month(), nowDay.Day(), 4, 59, 59, 59, time.Local)
	offsetDays := int(weekday - today.Weekday())
	absDay := int(math.Abs(float64(weekday - today.Weekday()))) //
	if offsetDays < 0 {
		offsetDays = 7 - absDay
	}
	targetDay := today.AddDate(0, 0, offsetDays)
	return targetDay.Sub(nowDay).Milliseconds()
}

func Now5OClock() time.Time {
	nowDay := time.Now()
	return time.Date(nowDay.Year(), nowDay.Month(), nowDay.Day(), 4, 59, 59, 59, time.Local)
}

func Now5OClockBeforeDays(days int) time.Time {
	nowDay := time.Now()
	now5Oclock := time.Date(nowDay.Year(), nowDay.Month(), nowDay.Day(), 4, 59, 59, 59, time.Local)
	return now5Oclock.AddDate(0, 0, -days)
}
