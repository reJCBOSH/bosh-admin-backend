package dao

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// CustomTime 自定义时间类型
type CustomTime time.Time

func (t CustomTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("\"\""), nil
	}
	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}

func (t *CustomTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) > 2 && data[0] == '"' {
		now, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
		*t = CustomTime(now)
		return err
	}
	return
}

func (t CustomTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tTime := time.Time(t)
	// 判断给定时间是否和默认零时间的时间戳相同
	if tTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tTime, nil
}

func (t *CustomTime) Scan(v any) error {
	value, ok := v.(time.Time)
	if ok {
		*t = CustomTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t CustomTime) String() string {
	return time.Time(t).Format(time.DateTime)
}

func (t CustomTime) ToTime() time.Time {
	tTime := time.Time(t)
	if tTime.IsZero() {
		return tTime
	}
	year, month, day := tTime.Date()
	return time.Date(year, month, day, tTime.Hour(), tTime.Minute(), tTime.Second(), 0, time.Local)
}
