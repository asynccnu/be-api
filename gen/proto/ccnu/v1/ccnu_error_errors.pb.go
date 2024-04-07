// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package ccnuv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsInvalidSidOrPwd(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CCNUErrorReason_INVALID_SID_OR_PWD.String() && e.Code == 401
}

func ErrorInvalidSidOrPwd(format string, args ...interface{}) *errors.Error {
	return errors.New(401, CCNUErrorReason_INVALID_SID_OR_PWD.String(), fmt.Sprintf(format, args...))
}

func IsNetworkToXkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == CCNUErrorReason_NETWORK_TO_XK_ERROR.String() && e.Code == 500
}

func ErrorNetworkToXkError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, CCNUErrorReason_NETWORK_TO_XK_ERROR.String(), fmt.Sprintf(format, args...))
}