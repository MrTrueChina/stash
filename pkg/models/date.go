package models

import "time"

// Date wraps a time.Time with a format of "YYYY-MM-DD"
type Date struct {
	time.Time
}

// 改版的时间格式，以原版代码中出现的精度最高的格式为准
const dateFormat = "2006-01-02 15:04:05.999999999-07:00"

// // 原版精确到日的时间
// const dateFormat = "2006-01-02"

func (d Date) String() string {
	return d.Format(dateFormat)
}

func NewDate(s string) Date {
	t, _ := time.Parse(dateFormat, s)
	return Date{t}
}
