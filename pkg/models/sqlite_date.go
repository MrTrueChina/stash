package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/stashapp/stash/pkg/utils"
)

// TODO - this should be moved to sqlite
type SQLiteDate struct {
	String string
	Valid  bool
}

// 改版的时间格式，以原版代码中出现的
const sqliteDateLayout = "2006-01-02 15:04:05.999999999-07:00"

// // 原版精确到日的时间
// const sqliteDateLayout = "2006-01-02"

// Scan implements the Scanner interface.
func (t *SQLiteDate) Scan(value interface{}) error {
	dateTime, ok := value.(time.Time)
	if !ok {
		t.String = ""
		t.Valid = false
		return nil
	}

	t.String = dateTime.Format(sqliteDateLayout)
	if t.String != "" && t.String != "0001-01-01" {
		t.Valid = true
	} else {
		t.Valid = false
	}
	return nil
}

// Value implements the driver Valuer interface.
func (t SQLiteDate) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}

	s := strings.TrimSpace(t.String)
	// handle empty string
	if s == "" {
		return "", nil
	}

	result, err := utils.ParseDateStringAsFormat(s, sqliteDateLayout)
	if err != nil {
		return nil, fmt.Errorf("converting sqlite date %q: %w", s, err)
	}
	return result, nil
}

func (t *SQLiteDate) StringPtr() *string {
	if t == nil || !t.Valid {
		return nil
	}

	vv := t.String
	return &vv
}

func (t *SQLiteDate) TimePtr() *time.Time {
	if t == nil || !t.Valid {
		return nil
	}

	ret, _ := time.Parse(sqliteDateLayout, t.String)
	return &ret
}

func (t *SQLiteDate) DatePtr() *Date {
	if t == nil || !t.Valid {
		return nil
	}

	ret := NewDate(t.String)
	return &ret
}
