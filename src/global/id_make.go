package global

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 生成订单id 年月日+4位顺序数字(例如：201802030001)
func MakeOrderId() (int64, error) {
	// 获取当前日期
	date := time.Now().Format("2006-01-02")
	dates := strings.Split(date, "-")

	ids := dates[0] + dates[1] + dates[2] + "0001"
	orderId, err := strconv.ParseInt(ids, 10, 64)
	return orderId, err
}

// 根据租期和是否试用得到订单的开始时间和结束时间
func GetStartEndTime(ifUse int8, rent int) (startTime, endTime int64) {
	// 获取当前日期的年月日
	year := time.Now().Year()
	months := time.Now().Month().String()
	day := time.Now().Day()
	years := strconv.Itoa(year)
	days := strconv.Itoa(day)
	h := time.Now().Hour()
	m := time.Now().Minute()
	s := time.Now().Second()
	hours := strconv.Itoa(h)
	switch len(hours) {
	case 1:
		hours = "0" + hours
	}
	minutes := strconv.Itoa(m)
	switch len(minutes) {
	case 1:
		minutes = "0" + minutes
	}
	seconds := strconv.Itoa(s)
	switch len(seconds) {
	case 1:
		seconds = "0" + seconds
	}
	// 转换月份
	months = ChangeMonth(months)
	var (
		newYear  int // 年
		newMonth int // 月
	)
	switch len(days) {
	case 1:
		days = "0" + days
	}
	if ifUse == 1 { // 试用(租期为7天)
		startTime = time.Now().Unix()
		endTime = startTime + SEVEN_DAYS
	} else if ifUse == 2 {
		if rent/12 == 0 { // 租期在一年以内
			newMonth, _ = strconv.Atoi(months)
			newYear, _ = strconv.Atoi(years)
			newMonth += rent
			if newMonth/12 > 0 {
				y, _ := strconv.Atoi(years)
				newYear = y + newMonth/12
				newMonth = newMonth % 12
			}
		} else { // 租期在一年以外
			newMonth, _ = strconv.Atoi(months)
			newYear, _ = strconv.Atoi(years)
			newYear += rent / 12
			newMonth += rent % 12
		}
		switch len(strconv.Itoa(newMonth)) {
		case 1:
			months = "0" + strconv.Itoa(newMonth)
		case 2:
			months = strconv.Itoa(newMonth)
		}
		toBeCharge := strconv.Itoa(newYear) + "-" + months + "-" + days + " " + hours + ":" + minutes + ":" + seconds
		// 获取时区
		loc, _ := time.LoadLocation("Local")
		theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", toBeCharge, loc)
		endTime = theTime.Unix()
		startTime = time.Now().Unix()
	}
	return startTime, endTime
}

// 转换星期
func ChangeWeek(week int) string {
	weekMap := map[int]string{
		1: "周一",
		2: "周二",
		3: "周三",
		4: "周四",
		5: "周五 ",
		6: "周六",
		7: "周日",
	}
	return weekMap[week]
}

// 转换月份
func ChangeMonth(months string) string {
	monthMap := map[string]string{
		"January":   "01",
		"February":  "02",
		"March":     "03",
		"April":     "04",
		"May":       "05",
		"June":      "06",
		"July":      "07",
		"August":    "08",
		"September": "09",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}
	return monthMap[months]
}

// 获取对应年月第一天0点时间戳
func FirstDayByMonth(year, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
}

// 获取对应年月最后一天时间戳
func LastDayByMonth(year, month int) time.Time {
	t := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	d, _ := time.ParseDuration("-24h")
	return t.Add(d)
}

// 获取今年对应月的第一天0点时间戳
func GetFirstDayByMonth(month int) int64 {
	times := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Local)
	return times.Unix()
}

// 获取今年第一天0点时间戳
func GetYearFirstDay() int64 {
	times := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local)
	return times.Unix()
}

// 获取当月第一天0时间戳
func GetMonthFirstDay() int64 {
	times := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	return times.Unix()
}

// 获取本周第一天0点时间戳
func FirstDayByCurrentWeek() time.Time {
	s := strconv.Itoa(-(int(time.Now().Weekday()-1) * 24)) + "h"
	d, _ := time.ParseDuration(s)
	m := time.Now().Add(d)
	return time.Date(m.Year(), m.Month(), m.Day(), 0, 0, 0, 0, time.Local)
}

// 获取上周第一天0点时间戳
func FirstDayByPreWeek() time.Time {
	d, _ := time.ParseDuration(strconv.Itoa(-7*24) + "h")
	return FirstDayByCurrentWeek().Add(d)
}

// 获取昨日0点时间戳
func PreDayByCurrent() time.Time {
	d, _ := time.ParseDuration("-24h")
	return CurrentDayByZeroHour().Add(d)
}

// 获取今日0点时间戳
func CurrentDayByZeroHour() time.Time {
	s := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", s, time.Local)
	return t
}

// 获取明日0点时间戳
func TomorrowByCurrent() time.Time {
	d, _ := time.ParseDuration("24h")
	return CurrentDayByZeroHour().Add(d)
}

// 时间戳转时间
func TimeStampTurnDate(ts int64) string {
	//格式化为字符串,tm为Time类型
	tm := time.Unix(ts, 0)
	return tm.Format("2006/01/02")
}

// float类型保留2位小数(四舍五入)
func FloatReserve2(src float64) float64 {
	temp, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", src), 64)
	return temp
}

// float类型保留4位小数(四舍五入)
func FloatReserve4(src float64) float64 {
	temp, _ := strconv.ParseFloat(fmt.Sprintf("%0.4f", src), 64)
	return temp
}
