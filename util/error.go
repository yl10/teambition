package util

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// XError error 扩展增加 错误码
type XError struct {
	code    int
	message string
}

// Error 获取错误内容
func (e *XError) Error() string {
	return e.message
}

// Code 获取错误编码
func (e *XError) Code() int {
	return e.code
}

// NewError 返回 XError 实例化指针
func NewError(msg string, err ...int) *XError {
	code := http.StatusBadRequest
	if len(err) > 0 {
		code = err[0]
	}

	message := "TB[" + strconv.Itoa(code) + "] " + msg
	fmt.Fprintf(os.Stderr, "%v [TB-%v] %v\n", time.Now().Format("2006/01/02 15:04:05"), code, msg)
	return &XError{code, message}
}

// TransErr2XErr 转换 error 为 XError
func TransErr2XErr(er error) *XError {
	return NewError(er.Error())
}

// TransErr2XErr 转换 error 为 XError
func TransErrs2XErr(er []error) *XError {
	var allErr []string
	for _, e := range er {
		allErr = append(allErr, e.Error())
	}

	return NewError(strings.Join(allErr, ", "))
}

func JoinErrs(er []error) error {
	var allErr []string
	for _, e := range er {
		allErr = append(allErr, e.Error())
	}

	return errors.New(strings.Join(allErr, ", "))
}
