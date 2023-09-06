package utils

import (
	"fmt"
	"time"
)

// Rails 的时间格式，Rails 是一个 Ruby 编写的 Web 服务框架。这里的时间格式是 YYYY-MM-DD hh:mm:ss 时区，这个格式似乎也是一个世界时间格式的标准，但我忘了是什么了，似乎是 RFC 的一种
const railsTimeLayout = "2006-01-02 15:04:05 MST"

func ParseDateStringAsTime(dateString string) (time.Time, error) {
	// https://stackoverflow.com/a/20234207 WTF?

	// 按照 RFC3339 格式进行转化，无错误则使用。
	// 【注意】这个方法不像 Java 那种有靠谱 Apache 支持的语言，它并不能应对所有的 RFC3339 支持的格式，所以下面还写了这么多的格式
	// RFC3339：YYYY-MM-DDTHH:mm:SS.ssssss±ZZ:ZZ，小写s用三个表示毫秒用六个则表示微秒（没有小数点和小s表示忽略毫秒），时间后面用加减号表示时区的时间偏移量(HH:mm)，单个字母表示时区（此时不需要加号）
	// T 可以用空格代替，即 YYYY-MM-DD HH:mm:SS.ssssss+ZZ:ZZ
	t, e := time.Parse(time.RFC3339, dateString)
	if e == nil {
		return t, nil
	}

	// YYYY-MM-DD 格式
	t, e = time.Parse("2006-01-02", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到秒、无时区格式
	t, e = time.Parse("2006-01-02 15:04:05", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到秒、加时区格式
	t, e = time.Parse("2006-01-02 15:04:05+07:00", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到秒、零时区格式
	t, e = time.Parse("2006-01-02 15:04:05Z", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到秒、减时区格式
	t, e = time.Parse("2006-01-02 15:04:05-07:00", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到纳秒、无时区格式
	t, e = time.Parse("2006-01-02 15:04:05.999999999", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到纳秒、零时区格式
	t, e = time.Parse("2006-01-02 15:04:05.999999999Z", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到纳秒、加时区格式
	t, e = time.Parse("2006-01-02 15:04:05.999999999+07:00", dateString)
	if e == nil {
		return t, nil
	}

	// 精确到纳秒、减时区格式
	t, e = time.Parse("2006-01-02 15:04:05.999999999-07:00", dateString)
	if e == nil {
		return t, nil
	}

	// 用 Rails 时间格式
	t, e = time.Parse(railsTimeLayout, dateString)
	if e == nil {
		return t, nil
	}

	return time.Time{}, fmt.Errorf("ParseDateStringAsTime failed: dateString <%s>", dateString)
}
