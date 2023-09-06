package models

import (
	"time"

	"github.com/stashapp/stash/pkg/utils"
)

// Date wraps a time.Time with a format of "YYYY-MM-DD"
type Date struct {
	time.Time
}

// 改版的时间格式，以原版代码中出现的那一串为准，但修改时区为东八区
const dateFormat = "2006-01-02 15:04:05.999999999+08:00"

// // 原版精确到日的时间
// const dateFormat = "2006-01-02"

func (d Date) String() string {
	return d.Format(dateFormat)
}

// ParseDate uses utils.ParseDateStringAsTime to parse a string into a date.
func ParseDate(s string) (Date, error) {
	ret, err := utils.ParseDateStringAsTime(s)
	if err != nil {
		return Date{}, err
	}
	return Date{Time: ret}, nil
}
