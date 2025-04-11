//go:build windows
// +build windows

package sotask

import (
	"errors"
	"syscall"

	ole "github.com/go-ole/go-ole"
)

var (
	ErrTargetUnsupported    = errors.New("error connecting to the Task Scheduler service: cannot connect to the XP or server 2003 computer")
	ErrConnectionFailure    = errors.New("error connecting to the Task Scheduler service: cannot connect to target computer")
	ErrInvalidPath          = errors.New(`path must start with root folder "\"`)
	ErrNoActions            = errors.New("definition must have at least one action")
	ErrInvalidPrinciple     = errors.New("both UserId and GroupId are defined for the principal; they are mutually exclusive")
	ErrRunningTaskCompleted = errors.New("the running task completed while it was getting parsed")
)

func getTaskSchedulerError(err error) error {
	errCode := getOLEErrorCode(err)
	switch errCode {
	case 50:
		return ErrTargetUnsupported
	case 0x80070032, 53:
		return ErrConnectionFailure
	default:
		return syscall.Errno(errCode)
	}
}

func getRunningTaskError(err error) error {
	errCode := getOLEErrorCode(err)
	if errCode == 0x8004130B {
		return ErrRunningTaskCompleted
	}

	return syscall.Errno(errCode)
}

func getOLEErrorCode(err error) uint32 {
	return err.(*ole.OleError).SubError().(ole.EXCEPINFO).SCODE()
}
