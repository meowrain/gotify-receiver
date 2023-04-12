package time

import (
	"fmt"
	"time"
)

type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (t *Time) String() string {
	return fmt.Sprintf("%s", time.Time(*t).Format(timeFormat))
}

// MarshalJSON on Json Time format Time field with %Y-%m-%d %H:%M:%S
func (t *Time) MarshalJSON() ([]byte, error) {
	// 重写time转换成json之后的格式
	var tmp = fmt.Sprintf("\"%s\"", t.String())
	return []byte(tmp), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	loc, _ := time.LoadLocation("Asia/Shanghai")
	rawT, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), loc)
	*t = Time(rawT)
	return err
}
