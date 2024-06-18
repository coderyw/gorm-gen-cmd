// Package model
// @Author: yinwei
// @File: time
// @Version: 1.0.0
// @Date: 2024/6/18 13:18

package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type GTime time.Time

func (g GTime) MarshalJSON() ([]byte, error) {
	format := "2006-01-02 15:04:05"

	formatted := fmt.Sprintf("\"%s\"", time.Time(g).Format(format))
	return []byte(formatted), nil
}

func (g GTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(g).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(g), nil
}

func (g *GTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*g = GTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
