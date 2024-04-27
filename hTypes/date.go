package hTypes

import (
	"reflect"
	"time"
)

type Date time.Time

var TimeConverter = func(value string) reflect.Value {
	if v, err := time.Parse("2006-01-02", value); err == nil {
		return reflect.ValueOf(v)
	}
	return reflect.Value{}
}

// Returns the default time.Time format of date
func (d *Date) Default() time.Time {
	t := time.Time(*d)
	return t
}

func (d *Date) ToString() string {
	return d.Default().Format("2006-01-02")
}
