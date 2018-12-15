package util

import (
	"fmt"
	"time"
)

type MyError struct {
	ms   string
	time time.Time
}

func (m *MyError) Error() string {
	return fmt.Sprintf("%s 。时间：%v", m.ms, m.time)
}

func NewMyError(s string) error {
	return &MyError{
		ms:   s,
		time: time.Now(),
	}

}
